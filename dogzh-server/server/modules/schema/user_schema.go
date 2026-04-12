package modules

import (
	"time"

	"github.com/atharvyadav96k/bus-safty/dogzh-server/modules/entites"
)

type UserSchema struct {
	ID         string        `json:'id'`
	Email      entites.Email `json:'email'`
	Password   string        `json:'-'`
	ImgUrl     string        `json:'img_url'`
	OrgId      int64         `json:'org_id'`
	Role       int           `json:'role'`
	RFIDId     int64         `json:'rfid`
	IsVerified bool          `json:'isVerified'`
	CreateAt   time.Time     `json:'create_at'`
	UpdateAt   time.Time     `json:'update_at'`
}
