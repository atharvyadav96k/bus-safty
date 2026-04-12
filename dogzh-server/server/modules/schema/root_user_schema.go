package modules

import (
	"time"

	"github.com/atharvyadav96k/bus-safty/dogzh-server/modules/entites"
)

type RootUserSchema struct {
	ID        string        `json:"id"`
	Email     entites.Email `json:"email"`
	OrgID     string        `json:"org_id"`
	Password  string        `json:"password"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
