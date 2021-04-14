package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"regexp"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func main() {

	pattern := flag.String("p", "site:google.com viagra", "google search pattern.")
	limit := flag.Int("min", 19, "minimal webhosts links limit for reporting")
	ignore := flag.String("i", "", "ignore regexp")
	verbose := flag.Bool("v", false, "verbose mode")

	flag.Parse()

	var re *regexp.Regexp

	if len(*ignore) > 0 {
		re = regexp.MustCompile(*ignore)
	}

	results, _ := googlesearch.Search(context.TODO(), *pattern, googlesearch.SearchOptions{Limit: 100, OverLimit: true})

	hosts := make(map[string]int)

	mainreturn := 0

	for _, u := range results {
		h, err := url.Parse(u.URL)
		if err == nil {
			if len(*ignore) > 0 {
				find := re.Find(([]byte)(h.Host))
				if len(find) == 0 {
					hosts[h.Host]++
				}
			} else {
				hosts[h.Host]++
			}
		}
	}

	for i, v := range hosts {
		if v > *limit {
			mainreturn++
			fmt.Printf("Please check webhost %s : %d Google links founds when searching \"%s\"\n", i, v, *pattern)
			if *verbose {
				for _, u := range results { // Would need caching, sub-optimal
					h, err := url.Parse(u.URL)
					if err == nil {
						if i == h.Host {
							fmt.Printf("%s [%s]\n", u.URL, u.Title)
						}
					}
				}
			}
		}
	}
	os.Exit(mainreturn)
}
