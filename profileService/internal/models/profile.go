package models

import (
	"encoding/json"
	"time"
)

type Profile struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Info        string `json:"info"`
	Items       []Item `json:"items"`
	LastUpdated time.Time
}

func (p *Profile) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

func ProfileFromJSON(profile []byte) (*Profile, error) {
	p := new(Profile)
	err := json.Unmarshal(profile, p)
	return p, err
}
