package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("--copy-file-names-empty.go--ver0.0.5")

	files, _ := filepath.Glob("*.mp4")
	if len(files) == 0 {
		files, _ = filepath.Glob("*.MP4")
	}
	if len(files) > 0 {
		os.Mkdir("CFNE", 0755 )
	}
	for i := range files {
		path := filepath.Join("CFNE", files[i]);
		//fmt.Println(files[i])
		//CFNE = Copy File Names Empty
		//f, err := os.Create("CFNE\\" + files[i])
		f, err := os.Create(path)
		check(err)
		defer f.Close()
	}

}
