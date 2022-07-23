package auth

import (
	"context"
	"log"

	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	oauth2models "github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

func initOauth() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	clientStore := store.NewClientStore()
	clientStore.Set(idvar, &oauth2models.Client{
		ID:     idvar,
		Secret: secretvar,
		Domain: domainvar,
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetPasswordAuthorizationHandler(func(ctx context.Context, clientID, username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})
	ginserver.SetUserAuthorizationHandler(userAuthorizeHandler)
	ginserver.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})
	ginserver.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)
}
