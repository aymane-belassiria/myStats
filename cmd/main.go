package main

import (
	"fmt"
	"log"
	"myStats/internals/runtime"
)

func main() {
	rt, err := runtime.NewDockerRuntime()
	if err != nil {
		log.Fatal(err)
	}
	containers, err := rt.ListContainers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Running Containers")
	for _, ctr := range containers {
		fmt.Printf("- %s - %s - %s - (%s)\n", ctr.Name, ctr.Image, ctr.ImageID, ctr.ID)
	}
}
