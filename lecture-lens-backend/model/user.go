package model

import "time"

type User struct {
	UserID    uint      `json:"userid" gorm:"primaryKey"`
	UserName  string    `json:"username"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
<<<<<<< HEAD
	UserID    uint      `json:"userid" gorm:"primaryKey"`
	UserName  string    `json:"username"`
	Email     string    `json:"email" gorm:"unique"`
}
=======
	UserID   uint   `json:"userid" gorm:"primaryKey"`
	UserName string `json:username`
	Email    string `json:"email" gorm:"unique"`
}
>>>>>>> e4d32e0 (update Auth)
