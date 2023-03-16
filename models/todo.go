package models

import "time"

type Todo struct {
	Id           int64
	Item         string
	Completed    bool
	Focused      bool
	Repeated     bool
	PostponeDate time.Time
}

func (t *Todo) GetPostponeDateAsString() string {
	return t.PostponeDate.Format("2006-02-01")
}

func (t *Todo) IsPostponed() bool {
	return t.PostponeDate.After(time.Now())
}
