package main

import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	bashCmd := exec.Command("bash", "dummy.sh")
	bashCmd.Stdout, bashCmd.Stderr = &stdout, &stderr
	err := bashCmd.Run()

	log.Printf("stdout: %s", &stdout)
	log.Printf("stderr: %s", &stderr)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
