package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/ngonzalezo/twitterGo/bd"
	"github.com/ngonzalezo/twitterGo/jwt"
	"github.com/ngonzalezo/twitterGo/models"
	"net/http"
	"time"
)

func Login(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi
	r.Status = 400
	fmt.Println("Ingresamos a login handler")
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		r.Message = "Usuario y/o contraseñas Invalidos " + err.Error()
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "El email del usuario es requerido "
		return r
	}
	fmt.Println("Vamoa a llamar a la funcion que intenta el loging")
	userData, existe := bd.IntentoLogin(t.Email, t.Password)

	if !existe {
		r.Message = "" + err.Error()
		return r
	}
	jwtKey, err := jwt.GeneroJWT(ctx, userData)

	if err != nil {
		r.Message = "Ocurrio un error al intentar generar el token correspondiente >" + err.Error()
		return r
	}

	resp := models.RespuestaLogin{
		TokenString: jwtKey,
	}

	token, err2 := json.Marshal(resp)

	if err2 != nil {
		r.Message = "Ocurrio un error al formatear el token a json >" + err2.Error()
		return r
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(time.Hour * 24),
	}

	cookieString := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  "*",
		},
	}

	fmt.Println(cookieString)

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res

	return r
}
