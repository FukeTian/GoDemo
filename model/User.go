package model

import "time"

type User struct {
	ID         uint      `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	CreateTime time.Time `json:"createtime"`
}
