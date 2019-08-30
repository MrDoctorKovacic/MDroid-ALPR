package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func alprImage(filename string) {
	log.Println("running ALPR on ", filename)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	out, err := exec.Command(fmt.Sprintf("docker run -it --rm -v %s/test/:/data:ro openalpr -j -c us h786poj.jpg", dir)).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
