package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func run(proj string, out io.Writer) error {
	if proj == "" {
		fmt.Errorf("Project diretory is required")
	}

	args := []string{"build", ".", "errors"}

	cmd := exec.Command("go", args...)
	cmd.Dir = proj

	if err := cmd.Run(); err != nil {
		return &stepErr{step: "go build", msg: "go build failed", cause: err}
	}

	_, err := fmt.Fprintln(out, "Go Build: SUCCESS")

	return err
}

func main() {
	proj := flag.String("p", "", "Project directory")
	flag.Parse()

	if err := run(*proj, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Hello World!")
}
