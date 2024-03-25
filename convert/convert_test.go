package convert

import (
	"fmt"
	"strings"
	"testing"
)

func TestCheckMi(t *testing.T) {
	filepath := "./sample/compress.msg"
	mime, err := checkMime(filepath)

	if err != nil {
		t.Error("error ", err)
	}

	fmt.Println("mime found", mime)

	if !strings.Contains(mime, "zip") {
		t.Error("ga nemu", mime)
	}

	filepath = "./sample/plain.msg"
	mime, err = checkMime(filepath)

	if err != nil {
		t.Error("error ", err)
	}

	fmt.Println("mime found", mime)

	if !strings.Contains(mime, "text") {
		t.Error("ga nemu", mime, filepath)
	}

}

func TestOpenFile(t *testing.T) {
	filePath := "./sample/compress.msg"
	res, err := OpenZip(filePath)

	if err != nil {
		t.Error("error setan ", err, len(res), res)
	}

	// if res != "koplak" {
	// 	fmt.Println(res)
	// 	t.Error("error", res)
	// }
}
