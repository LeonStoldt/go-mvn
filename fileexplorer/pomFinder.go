package fileexplorer

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
)

const pomRegEx = "pom\\.xml$"

func FindRecursivePomPaths(dir string) []string {
	var paths []string

	err := filepath.WalkDir(dir, func(file string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		pomRegExMatcher, e := regexp.Compile(pomRegEx)
		if e != nil {
			log.Fatal(e)
		}

		if !d.IsDir() && pomRegExMatcher.MatchString(file) {
			paths = append(paths, file)
		}
		return nil
	})

	if err != nil {
		return nil
	}

	for i := range paths { // TODO - for debug purposes
		fmt.Println(paths[i])
	}
	return paths
}
