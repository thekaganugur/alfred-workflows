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

	suggs := getSuggestions(link)

	for _, sugg := range suggs {
		wf.NewItem(sugg).Valid(true).Icon(icon).UID(sugg).Autocomplete(sugg).Arg(sugg)
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

	var suggs []string
	err = json.Unmarshal([]byte(body), &suggs)
	if err != nil {
		log.Fatal(err)
	}
	return suggs
}
