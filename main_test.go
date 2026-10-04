package main

import "testing"

func validEvent() TimingEvent {
	return TimingEvent{
		MeetID:    "sussex-2027",
		Event:     12,
		Heat:      3,
		Lane:      4,
		SwimmerID: "swimmer-123",
		Distance:  50,
		TimeMS:    27321,
	}
}

func TestTimingEventValid(t *testing.T) {
	tests := []struct {
		name  string
		event TimingEvent
		want  bool
	}{
		{name: "valid event", event: validEvent(), want: true},
		{name: "empty meet ID", event: TimingEvent{Event: 12, Heat: 3, Lane: 4, Distance: 50, TimeMS: 27321}, want: false},
		{name: "lane zero", event: TimingEvent{MeetID: "sussex-2027", Event: 12, Heat: 3, Lane: 0, Distance: 50, TimeMS: 27321}, want: false},
		{name: "lane eleven", event: TimingEvent{MeetID: "sussex-2027", Event: 12, Heat: 3, Lane: 11, Distance: 50, TimeMS: 27321}, want: false},
		{name: "negative time", event: TimingEvent{MeetID: "sussex-2027", Event: 12, Heat: 3, Lane: 4, Distance: 50, TimeMS: -1}, want: false},
		{name: "zero distance", event: TimingEvent{MeetID: "sussex-2027", Event: 12, Heat: 3, Lane: 4, Distance: 0, TimeMS: 27321}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.event.Valid()
			if got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
