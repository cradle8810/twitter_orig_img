package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type twit struct {
	fileName string
	html     []byte
	pictURLs []string
}

func main() {
	var page twit

	//for test
	if len(os.Args) != 2 {
		fmt.Println("for test:", os.Args[0], "fileName")
		os.Exit(2)
	}

	page.fileName = os.Args[1]

	e := page.loadFile()
	checkError(e)

	page.findOrigPict()
	page.renameLargeToOrig()
	page.printURLs()
}

//Load file
func (t *twit) loadFile() error {
	var err error
	t.html, err = ioutil.ReadFile(t.fileName)

	if err != nil {
		return err
	}

	return nil
}

//Parse HTML
//https://godoc.org/golang.org/x/net/html
func (t *twit) findOrigPict() error {
	r := bytes.NewReader(t.html)
	token := html.NewTokenizer(r)

	for {
		tt := token.Next()

		switch tt {
		//EOF
		case html.ErrorToken:
			return nil

		// Find the tag(s) "<meta  property="og:image" content="(HERE)">"
		case html.StartTagToken:
			tagName, _ := token.TagName()
			if string(tagName) == "meta" {
				key, val, _ := token.TagAttr()
				if (string(key) == "property") && (string(val) == "og:image") {
					_, val, _ = token.TagAttr()
					t.pictURLs = append(t.pictURLs, string(val))
				}
			}
		}
	}
}

func (t *twit) renameLargeToOrig() {
	for index, url := range t.pictURLs {
		t.pictURLs[index] = strings.Replace(url, ":large", ":orig", -1)
	}
}

func (t *twit) printURLs() {
	for _, url := range t.pictURLs {
		fmt.Println(url)
	}
}

func checkError(e error) {
	if e != nil {
		os.Stderr.WriteString(e.Error())
		os.Exit(1)
	}
}
