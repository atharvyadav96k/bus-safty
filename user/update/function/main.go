package org_register

import (
	"context"
	"errors"
	"net/http"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/user/update/applayer"
	"github.com/atharvyadav96k/gcp/common/entity"
	"github.com/atharvyadav96k/gcp/common/req"
	"github.com/atharvyadav96k/gcp/common/res"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
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

	if errs := entity.ValidateStruct(&user); len(errs) != 0 {
		res.BadRequest(w, errs)
		return
	}

	if err := app.CheckForDuplicate(context.Background(), "user", user); err != nil {
		res.BadRequest(w, []error{err})
		return
	}

	if err := app.StoreUpdate(context.Background(), "user", user.WhiteListedEmailID.String(), user); err != nil {
		res.NotFound(w, []error{err})
		return
	}
	res.Success(w, "Org updated successfully", user)
}
