package v1

import "time"

// LogEvent contains log event data
type LogEvent struct {
	ProcessId string    `bson:"process_id"`
	Log       string    `bson:"log"`
	Step      string    `bson:"step"`
	CreatedAt time.Time `bson:"created_at"`
}
