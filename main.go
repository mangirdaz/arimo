package main

import (
	"os"

	"github.com/jinzhu/configor"
	rm "github.com/mangirdaz/regman"
	containers "github.com/mangirdaz/regman/containers"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// Arimo - main structure for Arimo config
type Arimo struct {
	rm rm.Registry
}

func main() {

	log.Info("Starting ARIMO")

	config := containers.Config{}
	configor.Load(&config, "config.yaml")

	rm := containers.NewInstance(config)

	controller := Arimo{
		rm: rm,
	}

	for _, image := range config.ImageList {
		controller.rm.Copy(image)
	}

}
