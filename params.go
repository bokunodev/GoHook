package main

import "strings"

type Params []string

func (p Params) String() string {
	var sb strings.Builder
	for _, each := range p {
		sb.WriteString(each)
	}

	return sb.String()
}

func (p *Params) Get() interface{} {
	return []string(*p)
}

func (p *Params) Set(s string) error {
	*p = append(*p, s)
	return nil
}
