package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// batchMove moves files that match the formatstring and start/stop numbers into the given folder
func batchMove(fmtString string, folder string, startno int64, stopno int64) {
	if startno > stopno {
		log.Printf("`batchMove` warning: no files will be moved because startno `%d` is greater than stopno `%d`", startno, stopno)
	}
	err := os.Mkdir(folder, 0755)
	if err != nil {
		log.Fatalf("`batchMove` error creating folder `%s`\nError: %s", folder, err)
	}
	for i := startno; i <= stopno; i++ {
		groupName := fmt.Sprintf(fmtString, i)
		searchPath := groupName + "*"
		fileList, err := filepath.Glob(searchPath)
		if err != nil {
			log.Fatalf("`batchMove` error in obtaining file list. \nError: %s", err)
		}
		if len(fileList) == 0 {
			log.Printf("`batchMove` warning: no files returned for path `%s`", searchPath)
		}
		for _, file := range fileList {
			err := os.Rename(file, fmt.Sprintf("%s/%s", folder, file))
			if err != nil {
				log.Fatalf("`batchMove` error in moving file %s to %s\nError: %s", file, folder, err)
			}
		}
	}
}
