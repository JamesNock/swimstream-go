package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func timingEventsHandler(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()

    var event TimingEvent
    if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }
    if !event.Valid() {
        http.Error(w, "invalid timing event", http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusAccepted)
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

	http.HandleFunc("POST /timing-events", timingEventsHandler)

    http.ListenAndServe(":8080", nil)
}
