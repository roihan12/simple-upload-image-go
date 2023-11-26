package request

import (
	// "github.com/go-playground/validator/v10"
)
type PhotoInput struct {
	Title    string `json:"title" form:"title" validate:"max=225,min=1"`
	Caption  string `json:"caption" form:"caption" validate:"max=225,min=1"`
	PhotoURL string `json:"photo_url" validate:"max=225,min=1"`
	UserID   int    `json:"user_id" validate:"max=225,min=1"`
}

type PhotoUpdate struct {
	Title   string `form:"title" validate:"max=225,min=1"`
	Caption string `form:"caption" validate:"max=225,min=1"`
	PhotoURL string `json:"photo_url" validate:"max=225,min=1"`
}


// func (req *PhotoInput) Validate() error {
// 	validate := validator.New()

// 	err := validate.Struct(req)

// 	return err
// }

// func (req *PhotoUpdate) Validate() error {
// 	validate := validator.New()

// 	err := validate.Struct(req)

// 	return err
// }