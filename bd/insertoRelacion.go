package bd

import (
	"context"
	"github.com/ngonzalezo/twitterGo/models"
)

func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relacion")
	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
