package routers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/ngonzalezo/twitterGo/bd"
	"github.com/ngonzalezo/twitterGo/models"
)

func AltaRelacion(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {

	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parametro ID es obligatorio"
		return r
	}

	var t models.Relacion
	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		r.Message = "Ocurrrio un error al insertar relacion " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar la relacion"
		return r
	}
	r.Status = 200
	r.Message = "Alta de relacion ok"
	return r
}
