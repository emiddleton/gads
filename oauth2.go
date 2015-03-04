package gads

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"io/ioutil"
)

type AuthConfig struct {
	file         string             `json:"-"`
	OAuth2Config *oauth2.Config     `json:"oauth2.Config"`
	OAuth2Token  *oauth2.Token      `json:"oauth2.Token"`
	tokenSource  oauth2.TokenSource `json:"-"`
	Auth         Auth               `json:"gads.Auth"`
}

func NewCredentials(ctx oauth2.Context) (ac AuthConfig, err error) {
	data, err := ioutil.ReadFile(*configJson)
	if err != nil {
		return ac, err
	}
	if err := json.Unmarshal(data, &ac); err != nil {
		return ac, err
	}
	ac.file = *configJson
	ac.tokenSource = ac.OAuth2Config.TokenSource(ctx, ac.OAuth2Token)
	ac.Auth.Client = ac.OAuth2Config.Client(ctx, ac.OAuth2Token)
	return ac, err
}

// Save writes the contents of AuthConfig back to the JSON file it was
// loaded from.
func (c AuthConfig) Save() error {
	configData, err := json.MarshalIndent(&c, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.file, configData, 0600)
}

// Token implements oauth2.TokenSource interface and store updates to
// config file.
func (c AuthConfig) Token() (token *oauth2.Token, err error) {
	// use cached token
	if c.OAuth2Token.Valid() {
		return c.OAuth2Token, err
	}

	// get new token from tokens source and store
	c.OAuth2Token, err = c.tokenSource.Token()
	if err != nil {
		return nil, err
	}
	return token, c.Save()
}
