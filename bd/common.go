package bd

import (
	"os"

	"gitbub.com/fabiorodolfo84/gambituser/models"
	"gitbub.com/fabiorodolfo84/gambituser/secretm"
	"github.com/fabiorodolfo84/gambituser/models"
	"github.com/fabiorodolfo84/gambituser/secretm"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err := secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
