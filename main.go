package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

// Workflow starts here
func run() {
	var query string

	if args := wf.Args(); len(args) > 0 {
		query = args[0]
		query = url.QueryEscape(query)
	}

	icon := &aw.Icon{Value: "./icon.png"}
	link := fmt.Sprintf("https://ac.tureng.co/?t=%s&l=entr", query)

	sugs := getSuggestions(link)

	for _, sug := range sugs {
		wf.NewItem(sug).Valid(true).Icon(icon).UID(sug).Autocomplete(sug).Arg(sug)
	}

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}

func getSuggestions(link string) []string {
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var sugs []string
	err = json.Unmarshal([]byte(body), &sugs)
	if err != nil {
		log.Fatal(err)
	}
	return sugs
}
