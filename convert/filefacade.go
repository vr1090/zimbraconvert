package convert

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var EXT = ".eml"
var SEARCH_EXT = ".msg"

func GenerateNewEmlPath(msgpath, newfolder string) string {
	base := filepath.Base(msgpath)
	pathWithoutExt := strings.TrimSuffix(base, filepath.Ext(base)) + EXT
	newFilePath := filepath.Join(newfolder, pathWithoutExt)
	return newFilePath
}

/*
*
remove file
*
*/
func MoveFile(msgPath, newPath string) error {
	newDir := filepath.Dir(newPath)
	err := os.MkdirAll(newDir, 0755)

	if err != nil {
		return err
	}

	return CopyFile(msgPath, newPath)
}

func CopyFile(src, dst string) error {
	fmt.Println("anjing ", dst)
	// Open the source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create the destination file
	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Copy the content from source to destination
	_, err = destinationFile.ReadFrom(sourceFile)

	if err != nil {
		return err
	}

	// Flush the destination file's write buffer
	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func WriteFile(content bytes.Buffer, pathEml string, pathmsg string) error {
	dir := filepath.Dir(pathEml)
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		return err
	}

	err = os.WriteFile(pathEml, content.Bytes(), 0644)

	if err != nil {
		return err
	}

	return nil

}

func FindMsgFile(msgdir string) ([]string, error) {
	var result = make([]string, 0)

	err := filepath.Walk(msgdir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Size() != 0 {
			if filepath.Ext(path) == SEARCH_EXT {
				fmt.Println("nemu nih", path)
				result = append(result, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
