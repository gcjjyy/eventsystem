package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Choice ...
type Choice struct {
	Title      string `json:"title"`
	Connection string `json:"connection"`
}

// Say ...
type Say struct {
	Name    string   `json:"name"`
	Content string   `json:"content"`
	Choices []Choice `json:"choices"`
}

// Event ...
type Event struct {
	UUID  string `json:"uuid"`
	Event string `json:"event"`

	Say Say `json:"say"`

	Connection string `json:"connection"`
}

// ProcessEvent ...
func (s *Say) ProcessEvent() {
	fmt.Printf("[%v]\n%v\n", s.Name, s.Content)

	if len(s.Choices) > 0 {
		fmt.Println("\n[선택하세요]")
		for i, choice := range s.Choices {
			fmt.Printf(" - %v: %v\n", i+1, choice.Title)
		}
	}
}

// FindEventByUUID ...
func FindEventByUUID(events []Event, uuid string) *Event {
	for _, event := range events {
		if event.UUID == uuid {
			return &event
		}
	}

	return nil
}

func main() {
	jsonFile, _ := os.Open("opening.json")

	buf, _ := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()

	var events []Event
	var nextEvent string

	json.Unmarshal([]byte(buf), &events)

	currentEvent := &events[0]
	nextEvent = currentEvent.Connection
	for {
		switch strings.ToLower(currentEvent.Event) {
		case "say":
			currentEvent.Say.ProcessEvent()
		}

		var command int32
		fmt.Print("\n>> ")
		fmt.Scanf("%d", &command)

		if strings.ToLower(currentEvent.Event) == "say" &&
			len(currentEvent.Say.Choices) > 0 {
			nextEvent = currentEvent.Say.Choices[command-1].Connection
		}

		currentEvent = FindEventByUUID(events, nextEvent)
		if currentEvent == nil {
			break
		}
		nextEvent = currentEvent.Connection
	}
}
