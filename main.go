package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	bysize map[int64][]string //files with the same sizes
	excl   Extensions         //list of file extensions that will be excluded
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr,
			"dups is a program for finding duplicate files in one or more\n"+
				"directories. Use the -e option for every file extension to be\n"+
				"excluded from checking.\n"+
				"For more details: https://github.com/codehill/dups\n\n"+
				"Usage: dups [-e ext] dir1 [dir2...]\n")
	}

	bysize = make(map[int64][]string, 0)
	excl = Extensions{}
}

func main() {
	flag.Var(&excl, "e", "Exclude files with this extension")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		return
	}

	for _, arg := range args {
		if _, err := os.Open(arg); err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("%s does not exist\n", arg)
			} else if os.IsPermission(err) {
				fmt.Printf("%s permission denied\n", arg)
			} else {
				fmt.Println(err)
			}

			os.Exit(1)
		}
	}

	for _, arg := range args {
		if err := filepath.Walk(arg, walkfn); err != nil {
			fmt.Println(err)
		}
	}

	byhash := make(map[string][]string, 0)
	for _, size := range bysize {
		if len(size) == 1 {
			continue
		}

		for _, file := range size {
			hash, err := fileHash(file)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if _, ok := byhash[hash]; ok {
				byhash[hash] = append(byhash[hash], file)
			} else {
				byhash[hash] = []string{file}
			}
		}
	}

	for _, dups := range byhash {
		if len(dups) == 1 {
			continue
		}

		for _, f := range dups {
			fmt.Println(f)
		}

		fmt.Println()
	}
}
