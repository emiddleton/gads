package gads

import (
//	"encoding/xml"
//	"fmt"
)

type AdParamService struct {
	Auth
}

type AdParam struct {
}

func NewAdParamService(auth *Auth) *AdParamService {
	return &AdParamService{Auth: *auth}
}

// Query is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdParamService#get
//
func (s AdParamService) Get(selector Selector) (adParams []AdParam, err error) {
	return adParams, ERROR_NOT_YET_IMPLEMENTED
}
