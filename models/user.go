package models

import (
	"database/sql"
	"log"
	"time"
)

type UserConfig struct {
	Name string `json:"name"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (uc *UserConfig) Check(){
	log.Printf("Hi my name is %v and my password is %v", uc.Name, uc.Password)
}

func (uc *UserConfig) validate_name() (bool, string) {

	n := uc.Name

	if len(n) < 3 {
		return false, "Username must be at least 3 characters"
	}


	return true, ""
}

func (uc *UserConfig) validate_password() (bool, string) {

	if uc.Password != uc.ConfirmPassword {
		return false, "Passwords do not match"
	}
	if len(uc.Password) < 8 {
		return false, "Passwords must be at least 8 characters"
	}

	return true, ""
}

func (uc *UserConfig) Validate() (bool, []string) {
	errs := []string{}

	if success, err := uc.validate_name(); !success {
		errs = append(errs, err)
	}

	if success, err := uc.validate_password(); !success {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}


type user struct {
	ID int

	Name string `json:"name"`
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time

	db *sql.DB

}

func User(cfg UserConfig, db *sql.DB) *user {
	u := &user{
		Name:cfg.Name,
		Password:cfg.Password,
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
		db:db,
	}

	return u
}

func (u *user) Type() string {
	return "user"
}

func (u *user) Connect(db *sql.DB){
	u.db = db
}

func (u *user) Save() error {
	_, err := u.db.Query("INSERT INTO users (Name, Password, CreatedAt, UpdatedAt) VALUES (?, ?, ?, ?)", u.Name, u.Password, u.CreatedAt, u.UpdatedAt)
	return err
}
