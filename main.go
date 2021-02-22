package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

var (
	dotenv = flag.String("f", ".env", "dotenv `file` to load")
	must   = flag.Bool("must", false, "must load .env or fail.")
)

func main() {
	flag.Parse()
	if err := godotenv.Load(*dotenv); err != nil {
		log.Println(err)
		if *must {
			os.Exit(1)
		}
	}
	name, args := head(flag.Args())
	if name == "" {
		log.Fatal("no command to run")
	}
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		e, ok := err.(*exec.ExitError)
		if !ok {
			log.Fatal(err)
		}
		os.Exit(e.ExitCode())
	}
}

func head(s []string) (string, []string) {
	if len(s) == 0 {
		return "", nil
	}
	if len(s) == 1 {
		return s[0], nil
	}
	return s[0], s[1:]
}
