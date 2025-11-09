package main

import (
	"flag"
	"fmt"
	"log"
	"myStats/internals/runtime"
)

func main() {
	pidFlag := flag.String("PID", "", "Container ID to get the info of the container (optional)")
	flag.Parse()
	rt, err := runtime.NewDockerRuntime()
	if err != nil {
		log.Fatal(err)
	}
	if *pidFlag != "" {
		c, err := rt.GetContainerByPID(*pidFlag)
		if err != nil {
			log.Fatalf("Error getting PID for container %s: %v", *pidFlag, err)
		}
		fmt.Printf("- %s - %s - %s - (%s)\n", c.Name, c.Image, c.ImageID, c.ID)
		return
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
