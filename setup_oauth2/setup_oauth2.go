package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/emiddleton/gads"
	"github.com/toqueteos/webbrowser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
)

var (
	googleConfigJSON = flag.String("credentials_json", "./credentials.json", "API credentials from Google in JSON")
	newConfigJSON    = flag.String("new_config_json", "./config.json", "API credentials & tokens for gads in JSON")
)

func main() {
	data, err := ioutil.ReadFile(*googleConfigJSON)
	if err != nil {
		log.Panic(err)
	}
	conf, err := Oauth2ConfigFromJSON(data)
	if err != nil {
		log.Panic(err)
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Enter the code returned after authorising access: ")
	webbrowser.Open(url)

	// Use the authorization code that is pushed to the redirect URL.
	// NewTransportWithCode will do the handshake to retrieve
	// an access token and initiate a Transport that is
	// authorized and authenticated by the retrieved token.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Panic(err)
	}
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Panic(err)
	}

	ac := gads.AuthConfig{
		OAuth2Config: conf,
		OAuth2Token:  tok,
		Auth: gads.Auth{
			CustomerId:     "INSERT_YOUR_CLIENT_CUSTOMER_ID_HERE",
			DeveloperToken: "INSERT_YOUR_DEVELOPER_TOKEN_HERE",
			UserAgent:      "tests (Golang 1.4 github.com/emiddleton/gads)",
		},
	}
	configData, err := json.MarshalIndent(&ac, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(*newConfigJSON, configData, 0600)
	if err != nil {
		log.Panic(err)
	}
}

// Oauth2ConfigFromJSON returns an oauth2.Config setup for adwords api access from a
// "Client ID for native application" credential, or "Service Account" credential in
// JSON format.
func Oauth2ConfigFromJSON(jsonKey []byte) (oac *oauth2.Config, err error) {
	// try to load "Service Account" credential
	oac, err = google.ConfigFromJSON(jsonKey, "https://adwords.google.com/api/adwords")
	if err == nil {
		return oac, err
	}

	// fallback to "Client ID for native application" credential
	var c struct {
		Installed struct {
			ClientID     string   `json:"client_id"`
			ClientSecret string   `json:"client_secret"`
			RedirectURIs []string `json:"redirect_uris"`
			AuthURI      string   `json:"auth_uri"`
			TokenURI     string   `json:"token_uri"`
		} `json:"installed"`
	}
	if err := json.Unmarshal(jsonKey, &c); err != nil {
		return nil, err
	}
	if len(c.Installed.RedirectURIs) < 1 {
		return nil, errors.New("oauth2: missing redirect URL in the client_credentials.json")
	}
	return &oauth2.Config{
		ClientID:     c.Installed.ClientID,
		ClientSecret: c.Installed.ClientSecret,
		Scopes: []string{
			"https://adwords.google.com/api/adwords",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.Installed.AuthURI,
			TokenURL: c.Installed.TokenURI,
		},
		RedirectURL: "oob",
	}, nil
}
