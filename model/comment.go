package model

import (
	"time"
)

type Comment struct {
	Time    time.Time
	Content string
}
