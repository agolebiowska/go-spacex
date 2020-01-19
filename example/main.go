package main

import (
	"fmt"
	"github.com/agolebiowska/go-spacex/spacex"
	"log"
)

func main() {
	c := spacex.NewClient(nil)

	dragons, err := c.Dragons.ListAll(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SpaceX Dragon names:")
	for _, dragon := range dragons {
		fmt.Println(dragon.Name)
	}

	events, err := c.HistoricalEvents.ListAll(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SpaceX historical events info:")
	for _, event := range events {
		fmt.Println(fmt.Sprintf("%v - %v", event.Title, *event.Links.Wikipedia))
	}

	launches, err := c.Launches.ListAll(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, launch := range launches {
		if len(launch.Ships) > 0 {
			fmt.Println(fmt.Sprintf("%v", launch.Ships))
		}
	}
}
