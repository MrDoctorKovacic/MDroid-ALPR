package main

import (
	"fmt"
	"log"
	"os/exec"
)

func alprImage(filename string) {
	log.Println("running ALPR on ", filename)

	out, err := exec.Command(fmt.Sprintf("/usr/bin/alpr -j -c us %s", filename)).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
