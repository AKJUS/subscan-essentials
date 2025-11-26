package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Session struct {
	SessionId  uint             `gorm:"type:int;primaryKey;autoIncrement:false" json:"session_id"`
	Validators SessionValidator `gorm:"type:json;" json:"validators"`
}

type SessionValidator []string

func (j SessionValidator) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j SessionValidator) Marshal() []byte {
	b, _ := json.Marshal(j)
	return b
}

func (j *SessionValidator) Scan(src interface{}) error { return json.Unmarshal(src.([]byte), j) }
