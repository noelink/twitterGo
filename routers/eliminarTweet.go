package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/ngonzalezo/twitterGo/bd"
	"github.com/ngonzalezo/twitterGo/models"
)

func EliminarTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parametro ID es obligatorio"
		return r
	}

	err := bd.BorroTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrio un error al intentar borrar el tweet" + err.Error()
		return r
	}

	r.Message = "Eliminar Tweet OK !"
	r.Status = 200
	return r
}
