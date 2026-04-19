package org_register

import (
	"context"
	"errors"
	"net/http"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/org/register/applayer"
	"github.com/atharvyadav96k/gcp/common/res"
	"google.golang.org/api/iterator"
)

func GetAllOrg(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res.BadRequest(w, []error{errors.New("Method not allowed")})
		return
	}
	app := applayer.Init()
	defer app.Close()
	var orgs []database_models.Org
	iter := app.StoreDoc("org").Documents(context.Background())
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			res.BadRequest(w, []error{err})
			return
		}

		var o database_models.Org
		if err := doc.DataTo(&o); err != nil {
			res.BadRequest(w, []error{err})
			return
		}

		orgs = append(orgs, o)
	}
	res.Success(w, "All orgs", orgs)
}
