package model

import "time"

type Calc_Result struct {
	uuid string
	json string
	started_at time.Time
	completed_at time.Time
}