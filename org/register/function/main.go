package org_register

import (
	"context"
	"errors"
	"net/http"
	"time"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/org/register/applayer"
	"github.com/atharvyadav96k/gcp/common/entity"
	"github.com/atharvyadav96k/gcp/common/req"
	"github.com/atharvyadav96k/gcp/common/res"
)

func OrgRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res.BadRequest(w, []error{errors.New("method not allowed")})
		return
	}

	app := applayer.Init()
	defer app.Close()

	var org database_models.Org
	if err := req.ParseBody(r, &org); err != nil {
		res.BadRequest(w, []error{err})
		return
	}

	errs := entity.ValidateStruct(&org)
	if len(errs) != 0 {
		res.BadRequest(w, errs)
		return
	}

	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()

	ctx := context.Background()
	if err := app.CreateRecord(ctx, &org, nil); err != nil {
		res.BadRequest(w, []error{err})
		return
	}

	res.Created(w, "Organization registered successfully", org)
}
