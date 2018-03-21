package main

import (
	"log"
	"os"
	"path/filepath"
)

func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	log.Println("rm -f file")
	os.Remove("file")

	log.Println("rm -rf dir2")
	os.RemoveAll("dir2")

	log.Println("rm -rf dir1/*")
	removeContents("dir1")

	log.Println("[ -f none ] && echo none exists")
	if _, err := os.Stat("none"); err == nil {
		log.Println("none exists")
	}

	log.Println("[ ! -f none ] && echo none dones't exist")
	if _, err := os.Stat("none"); os.IsNotExist(err) {
		log.Println("none doesn't exist")
	}
}
