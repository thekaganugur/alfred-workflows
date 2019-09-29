package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	// Package is called aw

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

// Your workflow starts here
func run() {
	var query string

	// ----------------------------------------------------------------
	// Handle CLI arguments
	// ----------------------------------------------------------------

	// You should always use wf.Args() in Script Filters. It contains the
	// same as os.Args[1:], but the arguments are first parsed for AwGo's
	// magic actions (i.e. `workflow:*` to allow the user to easily open
	// the log or data/cache directory).
	if args := wf.Args(); len(args) > 0 {
		// If you're using "{query}" or "$1" (with quotes) in your
		// Script Filter, $1 will always be set, even if to an empty
		// string.
		// This guard serves mostly to prevent errors when run on
		// the command line.
		query = args[0]

		// go runApp()
	}

	// ----------------------------------------------------------------
	// Filter items based on user query
	// ----------------------------------------------------------------

	// if query != "" {
	//
	// 	res := wf.Filter(query)
	//
	// 	log.Printf("%d results match \"%s\"", len(res), query)
	//
	// 	for i, r := range res {
	// 		log.Printf("%02d. score=%0.1f sortkey=%s", i+1, r.Score, wf.Feedback.Keywords(i))
	// 	}
	// }

	link := fmt.Sprintf("https://ac.tureng.co/?t=%s&l=entr", query)
	keys := getKeys(link)

	for _, e := range keys {
		wf.NewItem(e).Valid(true).Var("URL", link).Arg(query)

	}

	// Add a "Script Filter" result

	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}

func getKeys(link string) []string {
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

	var keys []string
	err = json.Unmarshal([]byte(body), &keys)
	if err != nil {
		log.Fatal(err)
	}

	// keys = []string{"kitap", "kitap", "akciğeri", "kitap", "altlığı", "kitap", "armağan"}

	return keys
}

func runApp() {

	cmd := exec.Command("bash -c 'open -g /Applications/Tureng.app'")
	_, err := cmd.Output()

	if err != nil {
		// println(err.Error())
		return
	}

	// print(string(stdout))
}
