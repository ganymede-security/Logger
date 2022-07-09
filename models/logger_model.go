package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Log struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Date  string             `json:"date,omitempty" validate:"required"`
	Time  string             `json:"time,omitempty" validate:"required"`
	File  string             `json:"file,omitempty" validate:"required"`
	Level string             `json:"level,omitempty" validate:"required"`
	Msg   string             `json:"msg,omitempty" validate:"required"`
	Args  string             `json:"args,omitempty" validate:"required"`
}
