package operation

import (
	"context"
	"net/http"
	"time"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/gcp/app"
	"github.com/atharvyadav96k/gcp/common/res"
	"google.golang.org/api/iterator"
)

func CreateOrg(w http.ResponseWriter, app *app.App, org database_models.Org) {
	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()

	if err := app.StoreCreateWithId(context.Background(), "org", org.ContactEmail.String(), org); err != nil {
		res.BadRequest(w, []error{err})
		return
	}
	res.Created(w, "Org registered successfully", org)
}

func GetOrgs(w http.ResponseWriter, app *app.App) {
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
	res.Success(w, "Org's", orgs)
}

func Delete(w http.ResponseWriter, app *app.App, org database_models.Org) {
	if err := app.StoreDelete(context.Background(), "org", org.ContactEmail.String()); err != nil {
		res.NotFound(w, []error{err})
		return
	}

	res.Success(w, "Org deleted successfully", org)
}

func Update(w http.ResponseWriter, app *app.App, org database_models.Org) {
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
