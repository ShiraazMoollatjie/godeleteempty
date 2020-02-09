package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var rootDir = flag.String("rootDir", ".", "The root directory to start from")
var dryRun = flag.Bool("dryRun", true, "A dryrun prints out all the directories to delete instead of actually deleting them")

func deleteEmptyDirs(path string, info os.FileInfo, err error) error {
	if !info.IsDir() || path == *rootDir {
		return nil
	}

	d, err := os.Open(path)
	if err != nil {
		return err
	}
	defer d.Close()

	f, err := d.Readdir(0)
	if err != nil {
		return err
	}

	if len(f) == 0 {
		if *dryRun {
			log.Printf("%s", path)
		} else {
			err := os.Remove(path)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	flag.Parse()

	// TODO(shiraaz): This should be a depth first search for empty directories.
	filepath.Walk(*rootDir, deleteEmptyDirs)
}
