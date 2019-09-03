package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/MrDoctorKovacic/MDroid-Core/logging"
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
	if len(data.Result) == 0 {
		MainStatus.Log(logging.Warning(), "Plate has no results, not processing")
		return
	}

	makeRequest(&data.Result[0].Plate, &data.Result[0].Confidence)
	if len(data.Result[0].Candidiates) == 0 {
		MainStatus.Log(logging.OK(), "Plate has no candidates, stopping at initial values")
		return
	}

	for _, result := range data.Result[0].Candidiates {
		makeRequest(&result.Plate, &result.Confidence)
	}

	MainStatus.Log(logging.OK(), fmt.Sprintf("Plate has %d candidates, successfully processed.", len(data.Result[0].Candidiates)))
}

func makeRequest(plate *string, percent *float64) {
	url := fmt.Sprintf("%s/alpr/%s/", Hostname, *plate)
	var jsonStr = []byte(fmt.Sprintf(`{"percent":%f}`, *percent))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		MainStatus.Log(logging.Error(), fmt.Sprintf("Error sending post request: %s", err))
	}
	defer resp.Body.Close()
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
