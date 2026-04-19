package org_register

import (
	"context"
	"errors"
	"net/http"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/user/delete/applayer"
	"github.com/atharvyadav96k/gcp/common/req"
	"github.com/atharvyadav96k/gcp/common/res"
)

func UsersDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		res.BadRequest(w, []error{errors.New("Method not allowed")})
		return
	}
	app := applayer.Init()
	defer app.Close()

	var user database_models.User
	if err := req.ParseBody(r, &user); err != nil {
		res.BadRequest(w, []error{err})
		return
	}

	if err := app.StoreDelete(context.Background(), "user", user.WhiteListedEmailID.String()); err != nil {
		res.BadRequest(w, []error{err})
		return
	}

	res.Success(w, "User deleted successfully", nil)
}
