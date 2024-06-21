package dto

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type JobsDtoOutput struct {
	Id             string    `json:"id" validate:"required,uuid4"`
	Title          string    `json:"title" validate:"required,min=5"`
	Description    string    `json:"description"`
	Company        string    `json:"company" validate:"required"`
	JobOpeningDate time.Time `json:"job_opening_date" validate:"required"`
}
