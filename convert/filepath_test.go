package convert

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestFilePath(t *testing.T){
	var path = "/root/setan/koplak.msg"
	ext := ".eml"
	
	base := filepath.Base(path)
	
	if base != "koplak.msg"{
		t.Error("paan nih", base, filepath.Ext(base))
	}

	join := strings.TrimSuffix(base, filepath.Ext(base)) + ext


	if join != "koplak.eml"{
		t.Error("nah ini", join)
	}
}