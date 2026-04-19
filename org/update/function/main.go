package org_register

import (
	"context"
	"errors"
	"net/http"
	"time"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/org/update/applayer"
	"github.com/atharvyadav96k/gcp/common/entity"
	"github.com/atharvyadav96k/gcp/common/req"
	"github.com/atharvyadav96k/gcp/common/res"
)

func UpdateOrg(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		res.BadRequest(w, []error{errors.New("Method not allowed")})
		return
	}
	app := applayer.Init()
	defer app.Close()

	var org database_models.Org
	if err := req.ParseBody(r, &org); err != nil {
		res.BadRequest(w, []error{err})
		return
	}

	if errs := entity.ValidateStruct(&org); len(errs) != 0 {
		res.BadRequest(w, errs)
		return
	}

	if err := app.CheckForDuplicate(context.Background(), "org", org); err != nil {
		res.BadRequest(w, []error{err})
		return
	}

	org.UpdatedAt = time.Now()

	if err := app.StoreUpdate(context.Background(), "org", org.ContactEmail.String(), org); err != nil {
		res.NotFound(w, []error{err})
		return
	}
	res.Success(w, "Org updated successfully", org)
}
