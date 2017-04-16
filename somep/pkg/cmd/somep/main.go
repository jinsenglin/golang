package main

import (
    "fmt"
    _ "crypto/md5"
    _ "crypto/sha1"
    _ "encoding/base64"
    _ "encoding/csv"
    _ "encoding/json"
    _ "errors"
    _ "flag"
    _ "log"
    _ "math/rand"
    _ "net/url"
    _ "os"
    _ "os/exec"
    _ "os/signal"
    _ "regexp"
    _ "syscall"
    _ "time"
)

func main() {
	server()
	fmt.Println("main started")
}
