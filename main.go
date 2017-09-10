package main

import (
	"fmt"
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

	item, err := rm.GetImages()
	if err != nil {
		log.Error(err)
	}
	for _, img := range item {
		tags, err := rm.GetTags(img)
		if err != nil {
			log.Error(err)
		}
		for _, tag := range tags {
			log.Debugf("%s:%s", img, tag)
			dstImg := getDestinationImage(fmt.Sprintf("%s:%s", img, tag), config)
			rm.Copy(fmt.Sprintf("%s:%s", img, tag), dstImg)
		}
	}

}
