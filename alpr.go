package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func alprImage(filename string) {
	log.Println("running ALPR on ", filename)

	_, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	out, err := exec.Command(fmt.Sprintf("/usr/bin/alpr -j -c us h786poj.jpg")).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
