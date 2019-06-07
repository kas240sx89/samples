package models

import(
	"encoding/json"
	"time"
)

type Profile struct{
	Id string    `json:"id"`
	Email string `json:"email"`
	LastUpdated time.Time
}

func(p *Profile)ToJSON() ([]byte, error){
	return json.Marshal(p)
}

func ProfileFromJSON(profile []byte) ( *Profile, error ) {
	p := new(Profile)
	err := json.Unmarshal(profile, p)
	return p, err
}