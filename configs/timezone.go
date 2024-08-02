package configs

import (
	"log"
	"time"
)

func InitTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Printf("Failed to load location timezone: %v\n", err)
		log.Fatal(err)
	}

	time.Local = ict
}
