package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func alprImage(filename string) {
	log.Println("running ALPR on ", filename)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	out, err := exec.Command(fmt.Sprintf("/usr/bin/alpr -j -c us %s/%s", dir, filename)).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
