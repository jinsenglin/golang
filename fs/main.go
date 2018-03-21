package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	// Open the source file for reading
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	// Open the destination file for writing
	d, err := os.Create(dst)
	if err != nil {
		return err
	}

	// Copy the contents of the source file into the destination file
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}

	// Return any errors that result from closing the destination file
	// Will return nil if no errors occurred
	return d.Close()
}

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
