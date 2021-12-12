package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"hash"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	XGithubEventHeader     = "X-Github-Event"
	XHubSignature256Header = "X-Hub-Signature-256"
)

func main() {
	log.SetFlags(log.Lshortfile)

	var (
		address,
		command,
		secret,
		event,
		ref string
		params          = Params{}
		shutdownTimeout time.Duration
	)

	flag.StringVar(&address, "address", "localhost:8000", "address for server to linten on")
	flag.StringVar(&command, "command", "", "the command to execute")
	flag.DurationVar(&shutdownTimeout, "shutdown-timeout", 5*time.Second, "time limit, before hanging process get killed")
	flag.Var(&params, "params", "parameter to be passed to the command, could be used multiple times.")
	flag.StringVar(&secret, "secret", "", "Github webhook secret")
	flag.StringVar(&event, "event", "push", "Github event type")
	flag.StringVar(&ref, "ref", "refs/heads/main", "Github ref")
	flag.Parse()

	if command == "" {
		flag.PrintDefaults()
		return
	}

	if secret == "" {
		log.Println("-secret is required")
		return
	}

	hasher := func() hash.Hash {
		return hmac.New(sha256.New, []byte(secret))
	}

	var restart = make(chan struct{}, 1)
	var stop = make(chan struct{}, 1)

	go runner(restart, stop, command, params)

	server := http.Server{
		Handler: handler(restart, hasher, event, ref),
		Addr:    address,
	}

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
		<-sigCh
		close(restart)
		close(stop)
		server.Shutdown(context.Background())
	}()

	log.Println("server listening on", address)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func handler(restart chan<- struct{}, getHasher func() hash.Hash, event, ref string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		eventType := r.Header.Get(XGithubEventHeader)
		signature := r.Header.Get(XHubSignature256Header)

		signature = strings.TrimPrefix(signature, "sha256=")

		if eventType != event {
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Panicln(err)
		}

		hasher := getHasher()
		hasher.Write(body)

		sum := hasher.Sum(nil)
		hexSum := hex.EncodeToString(sum)
		if !hmac.Equal([]byte(hexSum), []byte(signature)) {
			log.Println("signature do not match")
			return
		}

		var data GithubWebhookPush
		if err = json.Unmarshal(body, &data); err != nil {
			log.Panicln(err)
		}

		if ref != data.Ref {
			return
		}

		log.Println("send restart signal")
		restart <- struct{}{}
	})
}

func runner(restart, stop <-chan struct{}, command string, params []string) {
theLoop:
	for {

		select {
		case <-stop:
			break theLoop
		default:
		}

		func() {
			log.Println("(re)starting process")

			cmd := exec.Command(command, params...)

			errFile, err := os.OpenFile("command.err", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				log.Panicln(err)
			}
			defer errFile.Close()

			outFile, err := os.OpenFile("command.out", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				log.Panicln(err)
			}
			defer outFile.Close()

			errPipe, err := cmd.StderrPipe()
			if err != nil {
				log.Panicln(err)
			}
			defer errPipe.Close()

			outPipe, err := cmd.StdoutPipe()
			if err != nil {
				log.Panicln(err)
			}
			defer outPipe.Close()

			go copier(outFile, outPipe)
			go copier(errFile, errPipe)

			if err := cmd.Start(); err != nil {
				log.Panicln(err)
			}

			// waiting for restart signanl
			<-restart
			log.Println("received restart signal")

			done := make(chan struct{}, 1)

			log.Println("terminating process")
			go func() {
				defer close(done)
				if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
					log.Panicln(err)
				}
				cmd.Wait()
			}()

			select {
			case <-time.After(5 * time.Second):
				log.Println("process timeout after 5secs, killed")
				if err := cmd.Process.Kill(); err != nil {
					log.Panicln(err)
				}
			case <-done:
				log.Println("process stoped")
			}
		}()
	}

	log.Println("runner exit")
}

func copier(dst io.WriteCloser, src io.ReadCloser) {
	_, err := io.Copy(dst, src)

	if err != nil && !errors.Is(err, os.ErrClosed) {
		log.Panicln(err)
	}
}
