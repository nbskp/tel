package main

import (
	_ "embed"

	"github.com/reiver/go-telnet"
)

//go:embed content.txt
var content []byte

func main() {
	err := telnet.ListenAndServe(":23", newHandler())
	if nil != err {
		panic(err)
	}
}

type handler struct{}

func (h *handler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	w.Write(content)
	for {
	}
}

func newHandler() *handler {
	return &handler{}
}
