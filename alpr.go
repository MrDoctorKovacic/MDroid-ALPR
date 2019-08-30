package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type alpr struct {
	Version          int     `json:"version"`
	Width            int     `json:"img_width"`
	Height           int     `json:"img_height"`
	ProcessingTimeMS float32 `json:"processing_time_ms"`
	Results          result  `json:"results"`
}

type result struct {
	Plate       string   `json:"plate"`
	Confidence  float64  `json:"confidence"`
	Candidiates []result `json:"candidates"`
}

func processResults(jsons []byte) {
	var data alpr
	if err := json.Unmarshal(jsons, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func alprImage(filename string) {
	log.Println("running ALPR on ", filename)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	out, err := exec.Command("/usr/bin/alpr", "-j", "-c", "us", fmt.Sprintf("%s/%s", dir, filename)).Output()
	if err != nil {
		log.Fatal(err)
	}

	processResults(out)
}
