package helpers

import (
	"time"

	common_auth "github.com/atharvyadav96k/bus-safty/dogzh-server/common/auth"
	"github.com/atharvyadav96k/bus-safty/dogzh-server/modules/entites"
	modules "github.com/atharvyadav96k/bus-safty/dogzh-server/modules/schema"
)

func CreateRootUser(org *modules.OrgSchema) (modules.RootUserSchema, error) {
	random_password, _ := common_auth.RandomPassword()
	password, err := common_auth.Hash(random_password)
	if err != nil {
		return modules.RootUserSchema{}, err
	}
	rootUser := modules.RootUserSchema{
		Email:     entites.NewRootUser(org.Code),
		OrgID:     org.ID,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return rootUser, nil
}
