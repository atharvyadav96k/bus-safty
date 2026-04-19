package org_register

import (
	"errors"
	"net/http"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/org/register/applayer"
	"github.com/atharvyadav96k/bus-safty/org/register/operation"
	"github.com/atharvyadav96k/gcp/common/entity"
	"github.com/atharvyadav96k/gcp/common/req"
	"github.com/atharvyadav96k/gcp/common/res"
)

func OrgRegister(w http.ResponseWriter, r *http.Request) {
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

	switch r.Method {
	case http.MethodPost:
		operation.CreateOrg(w, app, org)
		return
	case http.MethodPatch:
		operation.Update(w, app, org)
		return
	case http.MethodGet:
		operation.GetOrgs(w, app)
		return
	case http.MethodDelete:
		operation.Delete(w, app, org)
		return
	default:
		res.BadRequest(w, []error{errors.New("Method not allowed")})
	}

}
