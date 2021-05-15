package usecases

import (
	"github.com/carrot-systems/cs-session/src/core/domain"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	jwt.StandardClaims
	UserId    string `json:"user_id"`
	SessionId string `json:"session_id"`
}

func (i interactor) generateJWT(session domain.Session) (string, error) {

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * 24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		SessionId: session.ID,
		UserId:    session.UserId,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", domain.ErrUnmarshallingFailed
	}

	return tokenString, nil
}

func (i interactor) CreateSession(request *domain.SessionCreationRequest) (string, error) {
	userId, err := i.userClientGateway.CheckCredentials(request.User, request.Credentials)

	if err != nil {
		return "", domain.ErrBadUserPassword
	}

	session, err := i.sessionRepo.CreateSession(userId)

	if err != nil {
		return "", err
	}

	jwt, err := i.generateJWT(*session)

	if err != nil {
		return "", nil
	}

	return jwt, nil
}

func (i interactor) DeleteSession(id string) error {
	err := i.sessionRepo.DeleteSession(id)

	return err
}
