package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"sync"

	aw "github.com/deanishe/awgo"
)

var (
	wf        *aw.Workflow
	err       error
	home      string
	wd        string
	gitIgnore string
	rg        string
	out       bytes.Buffer
	stderr    bytes.Buffer
	cmd       *exec.Cmd
	icon      *aw.Icon
	fileView  string
)

func init() {
	wf = aw.New()

	icon = &aw.Icon{Value: "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/MagnifyingGlassIcon.icns"}

	home, err = os.UserHomeDir()
	wd, err = os.Getwd()
	gitIgnore = wd + "/ignore"
	cmd = exec.Command("/usr/local/bin/fd", "--hidden", "--follow", "--ignore-file", gitIgnore)
	cmd.Dir = home

	cfg := aw.NewConfig()
	fileView = cfg.Get("fileView")
}

func run() {
	stdoutIn, _ := cmd.StdoutPipe()
	// stderrIn, _ := cmd.StderrPipe()
	err = cmd.Start()
	if err != nil {
		log.Printf("cmd.Start() failed with '%s'\n", err)
	}

	// Read from stdoutIn, wg ensures that we finish
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		scanner := bufio.NewScanner(stdoutIn)
		for scanner.Scan() {
			path := scanner.Text()

			if fileView == "true" {
				wf.NewFileItem(path)
			} else {
				it := wf.NewItem(path)
				it.Valid(true).UID(path).Autocomplete(path).Arg("~/" + path).Icon(icon)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Println(err, stderr)
		}

		wg.Done()
	}()
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
	}

	// Filter items based on user query
	var query string

	if args := wf.Args(); len(args) > 0 {
		query = args[0]
	}

	if query != "" {
		wf.Filter(query)

		// log.Printf("%d results match \"%s\"", len(res), query)
		// for i, r := range res {
		// 	log.Printf("%02d. score=%0.1f sortkey=%s", i+1, r.Score, wf.Feedback.Keywords(i))
		// }

	}
	wf.WarnEmpty("No matching folders found", "Try a different query?")

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
