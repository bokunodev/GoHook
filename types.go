package main

import (
	"net/url"
	"time"
)

type URL struct {
	*url.URL
}

func (u *URL) MarshalText() ([]byte, error) {
	return []byte(u.URL.String()), nil
}

func (u *URL) UnmarshalText(text []byte) (err error) {
	u.URL, err = url.Parse(string(text))
	return err
}

type GithubMeta struct {
	SSHKeyFingerprints               SSHKeyFingerprints `json:"ssh_key_fingerprints"`
	Packages                         []string           `json:"packages"`
	Hooks                            []string           `json:"hooks"`
	Web                              []string           `json:"web"`
	API                              []string           `json:"api"`
	Git                              []string           `json:"git"`
	Pages                            []string           `json:"pages"`
	Importer                         []string           `json:"importer"`
	Actions                          []string           `json:"actions"`
	Dependabot                       []string           `json:"dependabot"`
	VerifiablePasswordAuthentication bool               `json:"verifiable_password_authentication"`
}

type SSHKeyFingerprints struct {
	Sha256Rsa     string `json:"SHA256_RSA"`
	Sha256Ecdsa   string `json:"SHA256_ECDSA"`
	Sha256Ed25519 string `json:"SHA256_ED25519"`
}

type GithubWebhookPush struct {
	Pusher     Pusher     `json:"pusher"`
	Ref        string     `json:"ref"`
	BaseRef    string     `json:"base_ref"`
	Before     string     `json:"before"`
	After      string     `json:"after"`
	Compare    string     `json:"compare"`
	Commits    []Commit   `json:"commits"`
	Sender     Sender     `json:"sender"`
	HeadCommit HeadCommit `json:"head_commit"`
	Repository Repository `json:"repository"`
	Forced     bool       `json:"forced"`
	Deleted    bool       `json:"deleted"`
	Created    bool       `json:"created"`
}

type Commit struct {
	Timestamp time.Time `json:"timestamp"`
	URL       URL       `json:"url"`
	Author    Author    `json:"author"`
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Added     []string  `json:"added"`
	Modified  []string  `json:"modified"`
	Removed   []string  `json:"removed"`
	Distinct  bool      `json:"distinct"`
}

type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Committer struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type HeadCommit struct {
	Timestamp time.Time `json:"timestamp"`
	Author    Author    `json:"author"`
	Committer Committer `json:"committer"`
	Message   string    `json:"message"`
	URL       string    `json:"url"`
	ID        string    `json:"id"`
	TreeID    string    `json:"tree_id"`
	Added     []string  `json:"added"`
	Removed   []string  `json:"removed"`
	Modified  []string  `json:"modified"`
	Distinct  bool      `json:"distinct"`
}

type Owner struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Login             string `json:"login"`
	Type              string `json:"type"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	ID                int    `json:"id"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Repository struct {
	UpdatedAt        time.Time `json:"updated_at"`
	License          string    `json:"license"`
	Homepage         string    `json:"homepage"`
	MirrorURL        string    `json:"mirror_url"`
	Description      string    `json:"description"`
	SvnURL           string    `json:"svn_url"`
	HTMLURL          string    `json:"html_url"`
	DefaultBranch    string    `json:"default_branch"`
	MasterBranch     string    `json:"master_branch"`
	URL              string    `json:"url"`
	ForksURL         string    `json:"forks_url"`
	KeysURL          string    `json:"keys_url"`
	CollaboratorsURL string    `json:"collaborators_url"`
	TeamsURL         string    `json:"teams_url"`
	HooksURL         string    `json:"hooks_url"`
	IssueEventsURL   string    `json:"issue_events_url"`
	EventsURL        string    `json:"events_url"`
	AssigneesURL     string    `json:"assignees_url"`
	FullName         string    `json:"full_name"`
	TagsURL          string    `json:"tags_url"`
	BlobsURL         string    `json:"blobs_url"`
	GitTagsURL       string    `json:"git_tags_url"`
	GitRefsURL       string    `json:"git_refs_url"`
	TreesURL         string    `json:"trees_url"`
	StatusesURL      string    `json:"statuses_url"`
	LanguagesURL     string    `json:"languages_url"`
	StargazersURL    string    `json:"stargazers_url"`
	ContributorsURL  string    `json:"contributors_url"`
	SubscribersURL   string    `json:"subscribers_url"`
	SubscriptionURL  string    `json:"subscription_url"`
	CommitsURL       string    `json:"commits_url"`
	GitCommitsURL    string    `json:"git_commits_url"`
	CommentsURL      string    `json:"comments_url"`
	IssueCommentURL  string    `json:"issue_comment_url"`
	ContentsURL      string    `json:"contents_url"`
	CompareURL       string    `json:"compare_url"`
	MergesURL        string    `json:"merges_url"`
	ArchiveURL       string    `json:"archive_url"`
	DownloadsURL     string    `json:"downloads_url"`
	IssuesURL        string    `json:"issues_url"`
	PullsURL         string    `json:"pulls_url"`
	MilestonesURL    string    `json:"milestones_url"`
	NotificationsURL string    `json:"notifications_url"`
	LabelsURL        string    `json:"labels_url"`
	ReleasesURL      string    `json:"releases_url"`
	DeploymentsURL   string    `json:"deployments_url"`
	Language         string    `json:"language"`
	Name             string    `json:"name"`
	NodeID           string    `json:"node_id"`
	GitURL           string    `json:"git_url"`
	SSHURL           string    `json:"ssh_url"`
	CloneURL         string    `json:"clone_url"`
	BranchesURL      string    `json:"branches_url"`
	Owner            Owner     `json:"owner"`
	ID               int       `json:"id"`
	Size             int       `json:"size"`
	Stargazers       int       `json:"stargazers"`
	CreatedAt        int       `json:"created_at"`
	Watchers         int       `json:"watchers"`
	OpenIssues       int       `json:"open_issues"`
	Forks            int       `json:"forks"`
	StargazersCount  int       `json:"stargazers_count"`
	OpenIssuesCount  int       `json:"open_issues_count"`
	ForksCount       int       `json:"forks_count"`
	PushedAt         int       `json:"pushed_at"`
	WatchersCount    int       `json:"watchers_count"`
	Fork             bool      `json:"fork"`
	Disabled         bool      `json:"disabled"`
	HasPages         bool      `json:"has_pages"`
	HasDownloads     bool      `json:"has_downloads"`
	HasProjects      bool      `json:"has_projects"`
	HasIssues        bool      `json:"has_issues"`
	Private          bool      `json:"private"`
	Archived         bool      `json:"archived"`
	HasWiki          bool      `json:"has_wiki"`
}

type Pusher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Sender struct {
	Login             string `json:"login"`
	Type              string `json:"type"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	ID                int    `json:"id"`
	SiteAdmin         bool   `json:"site_admin"`
}
