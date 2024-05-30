package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func listFiles(dir string, name string) bool {
	result := false
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, "/"+name) {
			//fmt.Println(path)
			result = true
		}
		return nil
	})

	return result
}

func searchCommandInPath(commandName string) (string, bool) {
	path, ok := os.LookupEnv("PATH")

	if ok {
		dirs := strings.Split(path, ":")
		for _, dir := range dirs {
			if listFiles(dir, commandName) {
				return dir + "/" + commandName, true
			}
		}
	}
	return "", false
}

func execFile(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func parsePath(path string) string {
	if strings.HasPrefix(path, "/") {
		return path
	}

	//It is a relative path
	currentPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	return filepath.Join(currentPath, path)
}
