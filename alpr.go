package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type alpr struct {
	Version          int       `json:"version"`
	Width            int       `json:"img_width"`
	Height           int       `json:"img_height"`
	ProcessingTimeMS float32   `json:"processing_time_ms"`
	Result           []results `json:"results"`
}

type results struct {
	Plate       string      `json:"plate"`
	Confidence  float64     `json:"confidence"`
	Candidiates []plateData `json:"candidates"`
}

type plateData struct {
	Plate      string  `json:"plate"`
	Confidence float64 `json:"confidence"`
}

func runALPR(filename *string) ([]byte, error) {
	log.Println("running ALPR on ", filename)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return exec.Command("/usr/bin/alpr", "-j", "-c", "us", fmt.Sprintf("%s/%s", dir, *filename)).Output()
}

func processResults(jsons *[]byte) (alpr, error) {
	var data alpr
	if err := json.Unmarshal(*jsons, &data); err != nil {
		return data, err
	}
	return data, nil
}

// Finally post results to MDroid-Core
func postResults(filename *string, data *alpr) {

}

func alprImage(filename string) {
	out, err := runALPR(&filename)
	if err != nil {
		log.Println(err)
		return
	}

	data, resultErr := processResults(&out)
	if resultErr != nil {
		log.Println(resultErr)
		return
	}

	postResults(&filename, &data)
}
