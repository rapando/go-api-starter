package entities

import "errors"

type User struct {
	UserID     int64  `json:"user_id"`
	FullName   string `json:"full_name"`
	UserTypeID int64  `json:"user_type_id"`
	Email      string `json:"email"`
	Active     bool   `json:"active"`
	Password   string `json:"password"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func (u *User) Validate() (err error) {
	if u.FullName == "" {
		return errors.New("Full Name is required")
	}
	if u.UserTypeID == 0 {
		return errors.New("User Type ID is required")
	}

	if u.Email == "" {
		return errors.New("Email is required")
	}

	if u.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}



