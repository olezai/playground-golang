package main

import (
	"log"
	"zipdemo/zipdemo"
)

func main() {
	f := "nicefile.txt"
	a := "nicearchive.zip"
	c := "This is some random text that will be zipped.\nHello World!\nThis is a demo file."

	err := zipdemo.CreateFile(f, c)
	if err != nil {
		log.Fatal(err)
	}

	err = zipdemo.ZipFile(a, f)
	if err != nil {
		log.Fatal(err)
	}
}
