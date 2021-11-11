package v1

// ProcessEvent contains process web socket event
type ProcessEvent struct {
	ProcessId string                 `bson:"process_id"`
	Data      map[string]interface{} `bson:"data"`
}
