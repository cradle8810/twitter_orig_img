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
	picts    map[string]string
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

		case html.StartTagToken:
			tagName, _ := token.TagName()
			if 0 == strings.Compare(string(tagName), "meta") {
				fmt.Println(">>meta")
			}
		}

	}

}

func checkError(e error) {
	if e != nil {
		os.Stderr.WriteString(e.Error())
		os.Exit(1)
	}
}
