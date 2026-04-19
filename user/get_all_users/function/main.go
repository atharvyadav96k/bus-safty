package org_register

import (
	"context"
	"errors"
	"net/http"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/user/get_all_users/applayer"
	"github.com/atharvyadav96k/gcp/common/res"
	"google.golang.org/api/iterator"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		res.BadRequest(w, []error{errors.New("Method not allowed")})
		return
	}
	app := applayer.Init()
	defer app.Close()
	defer app.Close()
	var users []database_models.User
	iter := app.StoreDoc("users").Documents(context.Background())
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

		var o database_models.User
		if err := doc.DataTo(&o); err != nil {
			res.BadRequest(w, []error{err})
			return
		}

		users = append(users, o)
	}
	res.Success(w, "All users", users)
}
