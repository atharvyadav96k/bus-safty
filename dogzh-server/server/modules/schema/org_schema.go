package modules

import (
	"time"

	"github.com/atharvyadav96k/bus-safty/dogzh-server/modules/entites"
)

type OrgSchema struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Code         string        `json:"code"`
	ContactEmail entites.Email `json:"contact_email"`
	LogoURL      string        `json:"logo_url"`
	RootUserID   int64         `json:"root_user_id"`
	SubStart     time.Time     `json:"sub_start"`
	SubEnd       time.Time     `json:"sub_end"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}
