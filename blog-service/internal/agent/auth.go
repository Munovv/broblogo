package agent

import (
	"bytes"
	"encoding/json"
	"github.com/Munovv/broblogo/blog-service/internal/model"
	"net/http"
)

type auth struct {
	host string
	port string
}

func (a *auth) VerifyUser(in model.AuthServiceRequest) (model.AuthServiceResponse, error) {
	var out model.AuthServiceResponse

	reqBody, err := json.Marshal(&in)
	if err != nil {
		return out, err
	}

	resp, err := http.Post(a.buildPath(verifyPrefix, tokenPrefix), jsonType, bytes.NewBuffer(reqBody))
	if err != nil {
		return out, err
	}

	json.NewDecoder(resp.Body).Decode(&out)

	return out, nil
}

func (a *auth) buildPath(elements ...string) string {
	path := "http://" + a.host + ":" + a.port
	for _, element := range elements {
		path += element
	}

	return path
}

func NewAuthAgent() *auth {
	return &auth{}
}
