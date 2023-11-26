package request

import (
	// "github.com/go-playground/validator/v10"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/models"
)

type RegisterInput struct {
	Username        string `json:"username" binding:"required" validate:"required,max=225,min=1"`
	Email           string `json:"email" binding:"required,email" validate:"required,max=225,min=1"`
	Password        string `json:"password" binding:"required" validate:"required,max=225,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"required,min=6,eqfield=Password"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email" validate:"required,min=1"`
	Password string `json:"password" binding:"required" validate:"required,max=225,min=6"`
}

type UpdateInput struct {
	Username        string `json:"username" binding:"required" validate:"required,max=225,min=1"`
	Email           string `json:"email" binding:"required,email" validate:"required,max=225,min=1"`
	Password        string `json:"password" binding:"required" validate:"required,max=225,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"required,min=6,eqfield=Password"`
	User            models.User
}

// func (req *RegisterInput) Validate() error {
// 	validate := validator.New()

// 	err := validate.Struct(req)

// 	return err
// }

// func (req *LoginInput) Validate() error {
// 	validate := validator.New()

// 	err := validate.Struct(req)

// 	return err
// }

// func (req *UpdateInput) Validate() error {
// 	validate := validator.New()

// 	err := validate.Struct(req)

// 	return err
// }
