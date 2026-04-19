package org_register

import (
	"context"
	"errors"
	"net/http"
	"time"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/user/delete/applayer"
	"github.com/atharvyadav96k/gcp/common/entity"
	"github.com/atharvyadav96k/gcp/common/req"
	"github.com/atharvyadav96k/gcp/common/res"
)

func BusRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res.BadRequest(w, []error{errors.New("Method not allowed")})
	}

	app := applayer.Init()
	defer app.Close()

	var bus database_models.Vehicle
	if err := req.ParseBody(r, &bus); err != nil {
		res.BadRequest(w, []error{err})
		return
	}

	if errs := entity.ValidateStruct(&bus); len(errs) != 0 {
		res.BadRequest(w, errs)
		return
	}

	bus.CreatedAt = time.Now()
	bus.UpdatedAt = time.Now()

	if err := app.StoreCreate(context.Background(), "bus", bus); err != nil {
		res.BadRequest(w, []error{err})
		return
	}
	res.Created(w, "Bus registered successfully", bus)
}
