package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/fabiorodolfo84/gambituser/awsgo"
	"github.com/fabiorodolfo84/gambituser/bd"
	"github.com/fabiorodolfo84/gambituser/models"
)

func main() {
	lambda.Start(ExecutarLambda)

}

func ExecutarLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAws()

	if !ValidoParametros() {
		fmt.Println("Erro nos parámetros. deve enviar 'SecretManager'")
		err := errors.New("erro nos parámetros deve enviar SecretManager")
		return event, err
	}

	var datos models.Signup
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)

		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Erro ao ler o segredo " + err.Error())
		return event, err
	}

	err = bd.Signup(datos)
	return event, err

}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
