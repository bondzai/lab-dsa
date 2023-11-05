package utils

import (
	"log"
	"math"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		if err := godotenv.Load(); err == nil {
			value = os.Getenv(key)
		} else {
			log.Printf("Error loading .env file: %s\n", err)
		}
	}

	if value == "" {
		return fallback
	}

	return value
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func TrackPerformance(taskName string) func() {
	timeStart := time.Now()
	return func() {
		timeEnd := time.Now()
		log.Printf("%s took %v to run.\n", taskName, timeEnd.Sub(timeStart))
	}
}

// Helper function to compare floating point numbers
func AlmostEqual(a, b float64) bool {
	epsilon := 0.00000001
	return math.Abs(a-b) <= epsilon
}
