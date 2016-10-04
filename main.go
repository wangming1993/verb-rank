package main

import (
	"fmt"
	"github.com/wangming1993/verb-rank/lib"
)

func init() {
	fmt.Println("VERB-RANK INIT...")
}

func main() {
	golang()

	//wc()
}

func golang() {
	repo := &lib.Repo{
		Language: "go",
		Limit:    10,
	}
	repo.Search()
}

func wc() {
	dir := lib.CLONE_PATH
	lib.ReadDir(dir)
}
