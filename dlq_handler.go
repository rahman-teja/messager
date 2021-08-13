package messager

type DeadLetterQueueMessage struct {
	Subject           string         `json:"subject"`
	Publisher         string         `json:"publisher"`
	Consumer          string         `json:"consumer"`
	Key               string         `json:"key"`
	Headers           MessageHeaders `json:"headers"`
	Message           string         `json:"message"`
	CausedBy          string         `json:"causedBy"`
	FailedConsumeDate string         `json:"failedConsumeDate"`
}
