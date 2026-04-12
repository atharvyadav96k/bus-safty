package services

import (
	"fmt"
	"net/http"

	"github.com/atharvyadav96k/bus-safty/dogzh-server/common/helpers"
	"github.com/atharvyadav96k/bus-safty/dogzh-server/common/response"
	modules "github.com/atharvyadav96k/bus-safty/dogzh-server/modules/schema"
)

func Org_Create(w http.ResponseWriter, r *http.Request) {
	org, err := modules.OrgDecoder(r)
	if err != nil {
		response.HttpResponseBadRequest(w, "Failed to prase data")
	}
	user, err := helpers.CreateRootUser(&org)
	fmt.Println(user)
	response.HttpResponseCreated(w, "org created", user)
}
