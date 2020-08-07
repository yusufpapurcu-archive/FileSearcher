package app

import (
	"os/exec"
	"strings"
)

// FileSearch is a main function
func FileSearch(dir, keyword string) []string {
	searcher := exec.Command("rg", "--files", dir, "-g", "*"+keyword+"*")
	res, err := searcher.Output()
	if err != nil && err.Error() == "exit status 1" {
		return []string{"Can't fınd anything"}
	}
	return strings.Split(string(res), "\n")
}

// KeywordSearch is a main function
func KeywordSearch(dir, keyword string) []string {
	searcher := exec.Command("rg", keyword, dir)
	res, err := searcher.Output()
	if err != nil && err.Error() == "exit status 1" {
		return []string{"Can't fınd anything"}
	}
	return strings.Split(string(res), "\n")
}
