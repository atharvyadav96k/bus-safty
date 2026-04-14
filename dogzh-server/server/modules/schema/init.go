package modules

import (
	"net/http"

	"github.com/atharvyadav96k/bus-safty/dogzh-server/common/request"
)

func OrgDecoder(r *http.Request) (OrgSchema, error) {
	return request.Decoder[OrgSchema](r)
}

func RootUserDecoder(r *http.Request) (RootUserSchema, error) {
	return request.Decoder[RootUserSchema](r)
}

func UserDecoder(r *http.Request) (UserSchema, error) {
	return request.Decoder[UserSchema](r)
}
