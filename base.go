package gads

import (
	"bytes"
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	//	"log"
)

var (
	configJson = flag.String("config_json", "./config.json", "API credentials")
)

type Auth struct {
	CustomerId     string
	DeveloperToken string
	UserAgent      string
	Testing        *testing.T
	Client         *http.Client
}

type AuthConfig struct {
	OAuthConfig *oauthConfig `json:"oauth.Config"`
	OAuthToken  *oauth.Token `json:"oauth.Token"`
	Auth        Auth         `json:"gads.Auth"`
}

// hack to properly unmarshal TokenCache
type oauthConfig struct {
	*oauth.Config
}

func (m *oauthConfig) UnmarshalJSON(data []byte) error {
	oc := oauth.Config{}
	_ = json.Unmarshal(data, &oc)
	m.Config = &oc
	config := struct {
		CacheFile string `json:"TokenCache"`
	}{}
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}
	m.TokenCache = oauth.CacheFile(config.CacheFile)
	return nil
}

func NewCredentials() (authConfig AuthConfig, err error) {
	data, err := ioutil.ReadFile(*configJson)
	if err != nil {
		return authConfig, err
	}

	if err := json.Unmarshal(data, &authConfig); err != nil {
		return authConfig, err
	}

	transport := &oauth.Transport{Config: authConfig.OAuthConfig.Config}
	authConfig.OAuthToken.Expiry = time.Now()
	transport.Token = authConfig.OAuthToken
	authConfig.Auth.Client = transport.Client()

	return authConfig, err
}

//
// Selector structs
//
type DateRange struct {
	Min string `xml:"min"`
	Max string `xml:"max"`
}

type Predicate struct {
	Field    string   `xml:"field"`
	Operator string   `xml:"operator"`
	Values   []string `xml:"values"`
}

type OrderBy struct {
	Field     string `xml:"field"`
	SortOrder string `xml:"sortOrder"`
}

type Paging struct {
	Offset int64 `xml:"startIndex"`
	Limit  int64 `xml:"numberResults"`
}

type Selector struct {
	XMLName    xml.Name
	Fields     []string    `xml:"fields"`
	Predicates []Predicate `xml:"predicates"`
	DateRange  *DateRange  `xml:"dateRange,omitempty"`
	Ordering   []OrderBy   `xml:"ordering"`
	Paging     *Paging     `xml:"paging,omitempty"`
}

// error parsers
func selectorError() (err error) {
	return err
}

func (a *Auth) Request(serviceUrl, action string, body interface{}) (respBody []byte, err error) {
	type soapReqHeader struct {
		UserAgent        string `xml:"userAgent"`
		DeveloperToken   string `xml:"developerToken"`
		ClientCustomerId string `xml:"clientCustomerId"`
	}

	type soapReqBody struct {
		Body interface{}
	}

	type soapReqEnvelope struct {
		XMLName xml.Name      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Header  soapReqHeader `xml:"https://adwords.google.com/api/adwords/cm/v201309 Header>RequestHeader"`
		Body    soapReqBody   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	}

	reqBody, err := xml.MarshalIndent(
		soapReqEnvelope{
			Header: soapReqHeader{
				UserAgent:        a.UserAgent,
				DeveloperToken:   a.DeveloperToken,
				ClientCustomerId: a.CustomerId,
			},
			Body: soapReqBody{body},
		},
		"  ", "  ")
	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest("POST", serviceUrl, bytes.NewReader(reqBody))
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("Accept", "multipart/*")
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	contentLength := fmt.Sprintf("%d", len(reqBody))
	req.Header.Add("Content-length", contentLength)
	req.Header.Add("SOAPAction", action)
	if a.Testing != nil {
		a.Testing.Logf("request ->\n%s\n%#v\n%s\n", req.URL.String(), req.Header, string(reqBody))
	}
	resp, err := a.Client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	if a.Testing != nil {
		a.Testing.Logf("respBody ->\n%s\n%s\n", string(respBody), resp.Status)
	}

	type soapRespHeader struct {
		RequestId    string `xml:"requestId"`
		ServiceName  string `xml:"serviceName"`
		MethodName   string `xml:"methodName"`
		Operations   int64  `xml:"operations"`
		ResponseTime int64  `xml:"responseTime"`
	}

	type soapRespBody struct {
		Response []byte `xml:",innerxml"`
	}

	soapResp := struct {
		XMLName xml.Name       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Header  soapRespHeader `xml:"https://adwords.google.com/api/adwords/cm/v201309 Header>RequestHeader"`
		Body    soapRespBody   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &soapResp)
	if err != nil {
		return respBody, err
	}
	if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 405 || resp.StatusCode == 500 {
		fault := Fault{}
		err = xml.Unmarshal(soapResp.Body.Response, &fault)
		if err != nil {
			return respBody, err
		}
		return soapResp.Body.Response, &fault.Errors
	}
	return soapResp.Body.Response, err
}
