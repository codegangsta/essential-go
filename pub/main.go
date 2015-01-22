package main

import (
	"log"
	"os"

	"github.com/codegangsta/essential-go-2/pub/pages"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments")
	}

	filename := os.Args[1]
	p, err := pages.NewPage(filename)
	if err != nil {
		log.Fatalln(err)
	}

	err = p.Render("layout.html", os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
