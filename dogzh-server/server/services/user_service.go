package services

import (
	"net/http"
	"time"

	"github.com/atharvyadav96k/bus-safty/dogzh-server/common/response"
	modules "github.com/atharvyadav96k/bus-safty/dogzh-server/modules/schema"
)

func User_Register(w http.ResponseWriter, r *http.Request) {
	user, err := modules.UserDecoder(r)
	if err != nil {
		response.HttpResponseBadRequest(w, "Invalid Body")
	}
	user.CreateAt = time.Now()
	user.CreateAt = time.Now()
	response.HttpResponseCreated(w, "user created", user)
}
