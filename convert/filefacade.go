package convert

import (
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

	return os.Rename(msgPath, newPath)
}

func WriteFile(content string,pathEml string, pathmsg string) error{
	dir := filepath.Dir(pathEml)
	err := os.MkdirAll(dir,0755)

	if err!= nil {
		return err
	}

	err = os.WriteFile(pathEml,[]byte(content),0644)

	if err != nil {
		return err
	}

	return os.Remove(pathmsg)	

}


func FindMsgFile(msgdir string) ([]string, error) {
	var result = make([]string,0)

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
		return nil,err
	}

	return result,nil
}
