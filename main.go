package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	p := "."
	t := time.Now()
	if len(os.Args) > 1 {
		p = strings.Join(os.Args[1:], " ")
	}

	st, sterr := os.Stat(p)

	var err error
	if sterr != nil {
		_, err = os.OpenFile(p, os.O_CREATE, os.ModePerm)
	} else if st.IsDir() {
		err = filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			return os.Chtimes(path, t, t)
		})
	} else {
		err = os.Chtimes(p, t, t)
	}

	if err != nil {
		log.Fatal(err)
	}
}
