package convert

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

/*
*
check the MIME of the file
*
*/
func checkMime(fpath string) (string, error) {
	file, err := os.Open(fpath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	//
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return "", err
	}

	// Use http.DetectContentType to determine the MIME type
	mimeType := http.DetectContentType(buffer)
	return mimeType, nil
}

/*
*
open zip file
*
*/
func OpenZip(filepath string) (bytes.Buffer, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return bytes.Buffer{}, err
	}

	defer file.Close()

	gzReader, err := gzip.NewReader(file)

	if err != nil {
		fmt.Println("gagal baca zip disini kah?", err)
		return bytes.Buffer{}, err
	}

	defer gzReader.Close()

	var buffer bytes.Buffer

	if _, err := io.Copy(&buffer,gzReader); err != nil {
		return bytes.Buffer{}, err
	}

	return buffer, nil
}

func HandlePath(wg *sync.WaitGroup, path string, resultDir string) {
	newFile := GenerateNewEmlPath(path, resultDir)
	mime, err := checkMime(path)
	defer wg.Done()
	fmt.Println("processing ", path)

	if err != nil {
		fmt.Println("error check mime", err)
		os.Exit(1)
	}

	if strings.Contains(mime, "zip") {
		var res bytes.Buffer
		res, err = OpenZip(path)

		if err != nil {
			fmt.Println("error gagal parse", err, path)
			return
		}

		err = WriteFile(res, newFile, path)

	} else {
		err = MoveFile(path, newFile)
	}

	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	fmt.Println("done processing", newFile)
}
