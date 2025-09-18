package main

import (
	"fmt"
	"grep/wget/pkg"
	"log"
	"os"
	"regexp"
)

func main() {

	args := os.Args
	dir := "out"
	re := regexp.MustCompile(`^(https?://)?([a-zA-Z0-9.-]+\.[a-zA-Z]{2,})(:\d+)?(/.*)?$`)
	if len(args) >= 1 {
		if !re.MatchString(args[1]) {
			log.Fatal("Некорректный url")
		}
		if len(args) == 3 {
			dir = args[2]
		}

		resp, err := pkg.Downloader(args[1], "index.html", dir)
		if err != nil {
			log.Fatal(err)
		}

		links, err := pkg.LinkExtr(resp)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(links)
	}

}
