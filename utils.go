package main

import (
	"fmt"
	"strings"

	containers "github.com/mangirdaz/regman/containers"
)

func getDestinationImage(image string, config containers.Config) string {
	//TODO: add logic if image does not have namespace (library tipe)
	imageRaw := strings.Split(image, "/")
	imageNew := fmt.Sprintf("%s/%s", getMapNamespace(imageRaw[0], config), imageRaw[1])
	return imageNew
}

func getMapNamespace(source string, config containers.Config) string {
	//TODO add check if namespaces exist
	for _, element := range config.NamespaceMap {
		if element.Source == source {
			return element.Destination
		}
	}
	return source
}
