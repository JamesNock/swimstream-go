package main

import (
	"encoding/json"
	"fmt"
)

type TimingEvent struct {
	MeetID    string `json:"meet_id"`
	Event     int    `json:"event"`
	Heat      int    `json:"heat"`
	Lane      int    `json:"lane"`
	SwimmerID string `json:"swimmer_id"`
	Distance  int    `json:"distance"`
	TimeMS    int    `json:"time_ms"`
}

func (e TimingEvent) Valid() bool {
	if e.MeetID == "" {
		return false
	}
	if e.Event <= 0 || e.Heat <= 0 {
		return false
	}
	if e.Lane < 1 || e.Lane > 10 {
		return false
	}
	if e.Distance <= 0 || e.TimeMS <= 0 {
		return false
	}

	return true
}

func main() {
	input := `{
        "meet_id": "sussex-2027",
        "event": 12,
        "heat": 3,
        "lane": 4,
        "swimmer_id": "swimmer-123",
        "distance": 50,
        "time_ms": 27321
    }`

	var event TimingEvent
	err := json.Unmarshal([]byte(input), &event)
	if err != nil {
		fmt.Println("Could not decode JSON:", err)
		return
	}

	fmt.Printf("%+v\n", event)
}
