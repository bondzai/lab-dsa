package main

import (
	"time"

	"github.com/introbond/go-ez-toolbox/toolbox"
)

func main() {
	mockData := map[string]interface{}{
		"name":    "Test Data",
		"boolean": true,
		"number":  123,
	}

	toolbox.PPrint(mockData)

	end := toolbox.TrackPerformance("Test Function")
	time.Sleep(2 * time.Second)
	end()
}
