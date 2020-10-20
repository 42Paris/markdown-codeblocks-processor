package main

import (
	"fmt"
	"flag"
	"log"
	"os/exec"
	"os"
	"io"
	"regexp"
	"bufio"
	"strings"
	"time"
	"path/filepath"
)

type Action struct {
	Name string
	Action func(string)
}

var actions []Action
var fileName string
var playbook string
var execute = false

func run_script(interpretor string, command [][]byte) {
	if interpretor == "" {
		interpretor = "cat"
	}
	cmd := exec.Command("/usr/bin/env", interpretor)
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	out := bufio.NewScanner(stdout)
	err := bufio.NewScanner(stderr)
	go func() {
		for out.Scan() {
			fmt.Printf("(out) %s\n", out.Text())
		}
	}()
	go func() {
		for err.Scan() {
			fmt.Printf("\033[31m(err)\033[0;0m %s\n", err.Text())
		}
	}()
	cmd.Start()
	go func() {
		defer stdin.Close()
		script := ""
		for _, b := range(command) {
			script = script + "\n" + string(b)
		}
		io.WriteString(stdin, script)
	}()
	cmd.Wait()
}

func mdProcessor(file []string, files []string, rootPath string) {
	var command [][]byte
	var interpretor string

	parsing_command := false
	for i, line := range file {
		tmp := strings.TrimSpace(line)
		indent := strings.Repeat(" ", len(files) * 2) + "|"
		if match, _ := regexp.MatchString("^```$", tmp); parsing_command && ! match {
			command = append(command, []byte(line))
		} else if match, _ = regexp.MatchString("#.*.md", tmp); match {
			r := regexp.MustCompile(`^.*\(`)
			r2 := regexp.MustCompile(`\).*$`)
			tmp = r.ReplaceAllString(tmp, "")
			tmp = r2.ReplaceAllString(tmp, "")
			fmt.Printf(indent + "%d " + line + "\n", i)
			fmt.Printf(indent + "Proceed File: %s\n-------------------\n", tmp)
			path := rootPath + tmp
			cycle := false
			for _, f := range(files) {
				if f == path {
					cycle = true
					break;
				}
			}
			if ! cycle {
				files = append(files, path)
				mdProcessor(readFile(path), files, filepath.Dir(path) + "/")
			} else {
				fmt.Printf(indent + "Cycle Dependencies\n")
				os.Exit(1)
			}
		} else if match, _ = regexp.MatchString("^```$", tmp); match && parsing_command {
			parsing_command = false
			if interpretor == "" {
				interpretor = "cat"
			}
			fmt.Printf(indent + "Script (%s):\n", interpretor)
			for j, line := range command {
				fmt.Printf(indent + "%d |%d| " + string(line) + "\n", i + j - len(command), j)
			}
			if execute {
				fmt.Printf(indent + "Executing. .  .\n")
				run_script(interpretor, command)
			}
		} else if match, _ = regexp.MatchString("^```.*$", tmp); match {
			interpretor = tmp[3:]
			parsing_command = true
		} else {
			fmt.Printf(indent + "%d " + line + "\n", i)
		}
	}
}

func readFile(path string) (fileContent []string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("No such file or directory:", path)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fileContent = append(fileContent, tmp)
	}
	return
}

func main() {
	const Pg_Name = "Markdown Codeblocks Processor"
	var files []string

	flag.StringVar(&fileName, "logins", "", "File that contains logins to apply\n(One per line)")
	flag.StringVar(&playbook, "playbook", "", "File that contains todo list\n (YAML)")
	flag.BoolVar(&execute, "execute", false, "Specify this flag to execute script code from the .md")
	flag.Parse()
	fmt.Println(Pg_Name)
	fmt.Println(time.Now())
	//ask_menu(playbook)
	mdProcessor(readFile(playbook), files, filepath.Dir(playbook) + "/")
}
