package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/MrDoctorKovacic/MDroid-Core/logging"
	"github.com/MrDoctorKovacic/MDroid-Core/settings"
)

// MainStatus will control logging and reporting of status / warnings / errors
var MainStatus = logging.NewStatus("Main")
var Hostname string

func main() {

	// Start with program arguments
	var (
		settingsFile string
	)
	flag.StringVar(&settingsFile, "settings-file", "", "File to recover the persistent settings.")
	flag.Parse()

	// Parse settings file
	settingsData, _ := settings.SetupSettings(settingsFile)

	// Log settings
	out, err := json.Marshal(settingsData)
	if err != nil {
		panic(err)
	}
	MainStatus.Log(logging.OK(), "Using settings: "+string(out))

	// Parse through config if found in settings file
	config, ok := settingsData["MDROID"]
	if ok {
		// Set up bluetooth
		var ok bool
		Hostname, ok = config["MDROID_HOST"]
		if !ok {
			//log.Fatal("No config found in settings file, not parsing through config")
		}
	} else {
		//log.Fatal("No config found in settings file, not parsing through config")
	}

	if len(os.Args) != 2 {
		log.Fatal("Please provide the directory to watch for new JPEGs")
	}

	startALPRWatch()
}
