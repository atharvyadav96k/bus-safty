package org_register

import (
	"context"
	"net/http"

	"errors"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/org/register/applayer"
	"github.com/atharvyadav96k/gcp/common/entity"
	"github.com/atharvyadav96k/gcp/common/req"
	"github.com/atharvyadav96k/gcp/common/res"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func OrgRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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
	errs := entity.ValidateStruct(&org)

	if len(errs) != 0 {
		res.BadRequest(w, errs)
		return
	}

	if err := app.StoreCreateWithId(context.Background(), "org", org.ContactEmail.String(), org); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			err = errors.New("org already exits with this email")
		}
		res.BadRequest(w, []error{err})
		return
	}
	res.Created(w, "User registered successfully", org)
}
