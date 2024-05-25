package linux

import "github.com/briheet/gofetch/fetch"

type (
	Command string
	linux   struct{}
)

func New() fetch.Fetcher {
	return &linux{}
}
