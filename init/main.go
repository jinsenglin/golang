package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func reboot() {
	binary, lookErr := exec.LookPath("init")
	if lookErr != nil {
		log.Fatalln(lookErr)
	}
	args := []string{"init", "6"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		log.Fatalln(execErr)
	}
}

func main() {
	log.Println("Restarting system ...")
	time.Sleep(1 * time.Second)
	reboot()
}
