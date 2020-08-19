package base

import "time"

type MapFile struct {
	ID int `json:"ID"`
	Name string `json:"name"`
	Tag string `json:"tag"`
	JSON string `json:"json"`
	Changed time.Time `json:"tag"`
}


