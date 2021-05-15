package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/carrot-systems/cs-session/src/core/domain"
	"github.com/carrot-systems/cs-session/src/core/usecases"
	discoveryClient "github.com/carrot-systems/csl-discovery-client"
)

type userClientGateway struct {
	discovery *discoveryClient.DiscoveryClient
}

type userResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (u userClientGateway) getUserId(host string, user string, credentials domain.Credentials) (string, error) {
	url := fmt.Sprintf("http://%s/internal/%s/id?password=%s", host, user, credentials.Password)
	data, err := httpGET(url, nil, nil)

	if err != nil {
		return "", domain.ErrUnableToReachUserServer
	}

	var response userResponse

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(data.Body)

	if err != nil {
		return "", domain.ErrByteReading
	}

	err = json.Unmarshal(buf.Bytes(), &response)

	if err != nil {
		return "", domain.ErrUnmarshallingFailed
	}

	if !response.Success || response.Data == "" {
		return "", domain.ErrBadUserPassword
	}

	return response.Data, nil
}

func (u userClientGateway) getUserServer() (string, error) {
	service, err := u.discovery.GetService("user")

	if err != nil {
		return "", domain.ErrUnableToFindUserServer
	}

	if service != nil && len(service) <= 0 {
		return "", domain.ErrUnableToFindUserServer
	}

	var userService = service[0]

	return userService.ExternalUrl, nil
}

func (u userClientGateway) CheckCredentials(user string, credentials domain.Credentials) (string, error) {
	userServer, err := u.getUserServer()

	if err != nil {
		return "", err
	}

	uuid, err := u.getUserId(userServer, user, credentials)

	if err != nil {
		return "", err
	}

	return uuid, nil
}

func NewUserClientGateway(dc *discoveryClient.DiscoveryClient) usecases.UserClientGateway {
	return &userClientGateway{
		discovery: dc,
	}
}
