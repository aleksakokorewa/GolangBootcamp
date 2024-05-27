package fs_comparer

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func getInfo(filename string) (fileinfo []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileName := scanner.Text()
		fileinfo = append(fileinfo, fileName)
	}
	err = scanner.Err()
	return fileinfo, err
}

func CompareInfo(oldSnapshot string, newSnapshot string) {
	newInfo, err := getInfo(newSnapshot)
	if err != nil {
		fmt.Println(err)
	}
	oldInfo, err := getInfo(oldSnapshot)
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range newInfo {
		if !slices.Contains(oldInfo, value) {
			fmt.Printf("ADDED %s\n", value)
		}
	}
	for _, value := range oldInfo {
		if !slices.Contains(newInfo, value) {
			fmt.Printf("REMOVED %s\n", value)
		}
	}
}
