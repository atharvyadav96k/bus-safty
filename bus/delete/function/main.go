package org_register

import (
	"context"
	"errors"
	"net/http"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/bus/delete/applayer"
	"github.com/atharvyadav96k/gcp/common/req"
	"github.com/atharvyadav96k/gcp/common/res"
)

func DeleteBus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		res.BadRequest(w, []error{errors.New("Method not allowed")})
		return
	}
	app := applayer.Init()
	defer app.Close()

	var org database_models.Id
	if err := req.ParseBody(r, &org); err != nil {
		res.BadRequest(w, []error{err})
		return
	}

	if err := app.StoreDelete(context.Background(), "bus", org.ID); err != nil {
		res.NotFound(w, []error{err})
		return
	}

	res.Success(w, "Org deleted successfully", org)
}
