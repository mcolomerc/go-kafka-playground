package data

import (
	"fmt"
	"log"
	"mcolomerc/kafka-playground/config"
	"mcolomerc/kafka-playground/utils"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	tasks = 10
)

type MockData struct {
	Keys   []string
	Values []string

	Config *config.Config
}

func New(config config.Config) *MockData {
	m := &MockData{}
	m.Config = &config
	return m
}

func (m *MockData) GetKey(i int) string {
	return m.Keys[i]
}

func (m *MockData) GetValue(i int) string {
	return m.Values[i]
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func (m *MockData) BuildData(nummessages int, size int) ([]string, []string) {
	m.Keys = make([]string, nummessages)
	m.Values = make([]string, nummessages)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer timeTrack(time.Now(), "Build keys...")
		for i := 0; i < nummessages; i++ {
			m.Keys[i] = uuid.New().String()
		}
	}()
	go func() {
		var values []string
		c := make(chan []string, tasks)
		elements := nummessages / tasks
		modElements := nummessages % tasks
		values = append(values, m.buildValues(modElements, size)...)
		for i := 0; i < tasks; i++ {
			go func() {
				c <- m.buildValues(elements, size)
			}()
		}
		// wait for goroutines to finish
		for i := 0; i < tasks; i++ {
			values = append(values, <-c...)
		}
		m.Values = values
		out := fmt.Sprintf(">> Build values ... %d", len(m.Values))

		defer wg.Done()
		defer timeTrack(time.Now(), out)
	}()
	wg.Wait()
	return m.Keys, m.Values
}

func (m *MockData) buildValues(nummessages int, size int) []string {
	values := make([]string, nummessages)
	for i := 0; i < nummessages; i++ {
		values[i] = utils.RandStringBytesMaskImprSrcUnsafe(size)
	}
	return values
}
