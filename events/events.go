package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Events struct {
	Events []Event
}

type Event struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	EntityID  string    `json:"entity_id"`
	Version   int64     `json:"version"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
	UserID    string    `json:"user_id"`
}

type EventProcesser interface {
	ProcessEvent(event Event)
}

func CatchUp(ep EventProcesser, pointer string) {
	// get all the events since the pointer
	eventsUrl := os.Getenv("EVENT_HUB_API_URL") + "/events?count=100&from=" + pointer
	fmt.Print("Getting events from ", eventsUrl, "\n")
	resp, err := http.Get(eventsUrl)

	if err != nil {
		fmt.Print("Could not retrieve events, trying again...\n")
		time.Sleep(3 * time.Second)
		CatchUp(ep, pointer)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	var data []Event
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	// process each event found
	for _, e := range data {
		ep.ProcessEvent(e)
	}

	var last string

	if len(data) == 100 {
		last = data[99].ID
		CatchUp(ep, last)
	} else if len(data) > 0 {
		last = data[len(data)-1].ID
		go Listen(ep, last)
	} else {
		last = pointer
		go Listen(ep, last)
	}
}

func Listen(ep EventProcesser, pointer string) {
	time.Sleep(3 * time.Second)
	CatchUp(ep, pointer)
}
