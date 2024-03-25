package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"vragantha.id/emlconvert/convert"
)

type Request struct {
	Wg     *sync.WaitGroup
	Path   string
	Folder string
}

func main() {
	var root string
	var result string
	var numWorker int
	flag.StringVar(&root, "search", "", "search folder path for msg files")
	flag.StringVar(&result, "result", "", "result path")
	flag.IntVar(&numWorker,"worker",0,"jumlah worker")

	flag.Parse()

	if root == "" || result == "" || numWorker == 0 {
		fmt.Println("parameter di benerin dulu")
		os.Exit(1)
	}

	workerChan := make(chan Request, 10)

	for i:=0; i<numWorker;i++{
		go runnerWorker(workerChan)
	}
	

	runner(root, result, workerChan)

}

func runnerWorker(channel <-chan Request) {
	for req := range channel {
		convert.HandlePath(req.Wg, req.Path, req.Folder)
	}
}

func runner(rootpath string, resultPath string, chanWorker chan<- Request) {
	result, err := convert.FindMsgFile(rootpath)

	if err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup

	for count, path := range result {
		folderPath := count / 10000
		wg.Add(1)
		pathEnd := filepath.Join(resultPath, strconv.Itoa(folderPath))
		req := Request{Wg: &wg, Path: path, Folder: pathEnd}
		chanWorker <- req
	}

	wg.Wait()
}
