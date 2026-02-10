package decision

import (
	"time"

	"github.com/google/uuid"
)

func New(action, entity string) Decision {
	return Decision{
		ID:        uuid.NewString(),
		Action:    action,
		Entity:    entity,
		Reason:    "auto-recorded",
		Timestamp: time.Now(),
	}
}
