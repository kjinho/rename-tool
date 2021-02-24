package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func zipFolder(folder string, zipfilename string) error {
	zipfile, err := os.Create(zipfilename)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	fileinfo, err := os.Stat(folder)
	if err != nil {
		return err
	}
	if !fileinfo.IsDir() {
		return fmt.Errorf("`%s` is not a folder (cannot zip)", folder)
	}

	baseDir := filepath.Base(folder)
	walkFunc := func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fileinfo, err := info.Info()
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(fileinfo)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, folder))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	}

	filepath.WalkDir(folder, walkFunc)

	return nil
}

func genZipFilename(folder string) string {
	if strings.HasSuffix(folder, "/") {
		return folder[0:len(folder)-1] + ".zip"
	}
	return folder + ".zip"
}
