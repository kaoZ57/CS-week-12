package Create_file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
) 

func readtxt() []os.FileInfo {
	files, err := ioutil.ReadDir("./TXT/")
	if err != nil {
		log.Fatal(err)
	}

	return files
}
