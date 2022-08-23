package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "No urls detected. Hint: cat urls.txt | unhttpx")
		os.Exit(1)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		host, err := getHost(s.Text())
		if err != nil {
			log.Println("Error parsing url: ", err)
			return
		}
		fmt.Println(host)
	}
}

// gethost() extracts the hostname from a URL and returns it
func getHost(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	return u.Hostname(), nil
}
