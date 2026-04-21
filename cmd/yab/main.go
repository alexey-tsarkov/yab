package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/alexey-tsarkov/yab"
)

const TokenEnvName = "YANDEX_BOOKS_TOKEN"

var version = "dev"

var audios, books, comics yab.UUIDValues

func init() {
	flag.Var(&audios, "a", "audio book identifier")
	flag.Var(&audios, "audio", "audio book identifier (same as -a)")
	flag.Var(&audios, "audiobook", "audio book identifier (same as -a)")
	flag.Var(&books, "b", "book identifier")
	flag.Var(&books, "book", "book identifier (same as -b)")
	flag.Var(&comics, "c", "comic book identifier")
	flag.Var(&comics, "comic", "comic book identifier (same as -c)")
	flag.Var(&comics, "comicbook", "comic book identifier (same as -c)")

	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Tiny Yandex Books downloader (version %s)\n", version)
		fmt.Fprintf(w, "Usage of %s:\n", filepath.Base(os.Args[0]))
		fmt.Fprintln(w, "  -a|audio|audiobook UUID")
		fmt.Fprintln(w, "  -b|book UUID")
		fmt.Fprintln(w, "  -c|comic|comicbook UUID")
	}
}

func main() {
	flag.Parse()
	for _, b := range flag.Args() {
		flag.Set("book", b)
	}

	if len(audios) == 0 && len(books) == 0 && len(comics) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	token := os.Getenv(TokenEnvName)
	if len(token) == 0 {
		fmt.Fprintf(os.Stderr, "env variable not set: %s\n", TokenEnvName)
		os.Exit(2)
	}

	client := yab.NewClient(nil, token)
	if err := execute(context.Background(), client); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func execute(ctx context.Context, _ *yab.Client) error {
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	var wg sync.WaitGroup
	for _, uuid := range audios {
		wg.Go(func() {
			fmt.Printf("audiobook [%s]\n", uuid)
		})
	}
	for _, uuid := range books {
		wg.Go(func() {
			fmt.Printf("book [%s]\n", uuid)
		})
	}
	for _, uuid := range comics {
		wg.Go(func() {
			fmt.Printf("comicbook [%s]\n", uuid)
		})
	}
	wg.Wait()

	return context.Cause(ctx)
}
