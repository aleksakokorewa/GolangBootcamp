package fscomparer

import (
	"bufio"
	"fmt"
	"os"
)

func FSCompare(OldFile string, NewFile string) error {
	OldInfo, err := os.Open(OldFile)
	if err != nil {
		fmt.Println("Error opening file")
		return err
	}
	NewInfo, err := os.Open(NewFile)
	if err != nil {
		fmt.Println("Error opening file")
		return err
	}
	defer OldInfo.Close()
	defer NewInfo.Close()
	oldScanner := bufio.NewScanner(OldInfo)
	newScanner := bufio.NewScanner(NewInfo)
	var oldInfo = make(map[string]bool)
	var newInfo = make(map[string]bool)
	for oldScanner.Scan() {
		oldInfo[oldScanner.Text()] = true
	}
	for newScanner.Scan() {
		newInfo[newScanner.Text()] = true
	}
	for info := range oldInfo {
		if _, ok := newInfo[info]; !ok {
			fmt.Printf("ADDED %s\n", info)
		}
	}
	for info := range newInfo {
		if _, ok := oldInfo[info]; !ok {
			fmt.Printf("REMOVED %s\n", info)
		}
	}
	return nil
}
