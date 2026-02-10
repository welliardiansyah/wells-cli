package decision

import "time"

type Decision struct {
	ID        string    `json:"id"`
	Action    string    `json:"action"`
	Entity    string    `json:"entity"`
	Reason    string    `json:"reason"`
	Timestamp time.Time `json:"timestamp"`
}
