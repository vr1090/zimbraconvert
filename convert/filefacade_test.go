package convert

import "testing"

func TestGenerateNewEmlPath(t *testing.T) {
	filepath := "setan/gundul/koplak.msg"
	newPath:="babi"

	res := GenerateNewEmlPath(filepath,newPath)

	if res != "babi/koplak.eml"{
		t.Error("error nih", res)
	}
}


func TestMoveFile(t *testing.T){
	oldfile := "./sample/simple.msg"
	newfile := "./sample/generate/simple.eml"

	t.Skip("skip dulu, prepare filenya dulu")
	err := MoveFile(oldfile,newfile)

	if err != nil {
		t.Error("error", err)
	}
}

func TestFindMsgFile(t *testing.T){
	folder := "./sample/"

	res, err := FindMsgFile(folder)

	if err != nil {
		t.Error(err)
	}

	if len(res) == 0 {
		t.Error("hasilnya apaan?",res)
	}
}
