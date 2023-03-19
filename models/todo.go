package models

import "time"

type Todo struct {
	Id            int64
	Item          string
	Focused       bool
	Repeated      bool
	PostponeDate  time.Time
	CompletedDate *time.Time
}
