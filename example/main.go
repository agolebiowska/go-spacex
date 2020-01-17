package main

import (
	"fmt"
	"github.com/agolebiowska/go-spacex/spacex"
	"log"
)

func main() {
	c := spacex.NewClient(nil)

	dragons, err := c.Dragons.ListAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SpaceX Dragon names:")
	for _, dragon := range dragons {
		fmt.Println(dragon.Name)
	}

	events, err := c.HistoricalEvents.ListAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SpaceX historical events info:")
	for _, event := range events {
		fmt.Println(fmt.Sprintf("%v - %v", event.Title, event.Links.Wikipedia))
	}
}