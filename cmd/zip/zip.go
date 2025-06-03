package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 圧縮：指定されたファイルを zip ファイルとして保存する
func zipFile(sourceFile, zipFileName string) error {
	zipfile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	zipWriter := zip.NewWriter(zipfile)
	defer zipWriter.Close()

	fileToZip, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = filepath.Base(sourceFile)
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}

// 解凍：指定された zip ファイルを展開する
func unzip(zipFile, outputDir string) error {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(outputDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		inFile, err := f.Open()
		if err != nil {
			return err
		}
		defer inFile.Close()

		outFile, err := os.Create(fpath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, inFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// 圧縮するファイル
	source := "example.txt"
	zipName := "example.zip"
	unzipDir := "unzipped_output"

	// 圧縮処理
	fmt.Println("Compressing...")
	if err := zipFile(source, zipName); err != nil {
		fmt.Println("Error zipping file:", err)
		return
	}
	fmt.Println("Zipped:", zipName)

	// 解凍処理
	fmt.Println("Unzipping...")
	if err := unzip(zipName, unzipDir); err != nil {
		fmt.Println("Error unzipping file:", err)
		return
	}
	fmt.Println("Unzipped to:", unzipDir)
}
