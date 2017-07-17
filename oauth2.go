package gads

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type AuthConfig struct {
	file         string             `json:"-"`
	OAuth2Config *oauth2.Config     `json:"oauth2.Config"`
	OAuth2Token  *oauth2.Token      `json:"oauth2.Token"`
	tokenSource  oauth2.TokenSource `json:"-"`
	Auth         Auth               `json:"gads.Auth"`
}

type OAuthConfigArgs struct {
	ClientID     string
	ClientSecret string
}

type OAuthTokenArgs struct {
	AccessToken  string
	RefreshToken string
}

type Credentials struct {
	Config OAuthConfigArgs
	Token  OAuthTokenArgs
	Auth   Auth
}

func NewCredentialsFromFile(pathToFile string, ctx context.Context) (ac AuthConfig, err error) {
	data, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return ac, err
	}
	if err := json.Unmarshal(data, &ac); err != nil {
		return ac, err
	}
	ac.file = pathToFile
	ac.tokenSource = ac.OAuth2Config.TokenSource(ctx, ac.OAuth2Token)
	ac.Auth.Client = ac.OAuth2Config.Client(ctx, ac.OAuth2Token)
	return ac, err
}

func NewCredentialsFromParams(creds Credentials) (config AuthConfig, err error) {
	var gcfg AuthConfig

	expiresAt, _ := time.Parse(time.RFC3339, "2015-07-28T14:51:53.543430418-04:00")

	// Create a token
	gcfg.OAuth2Token = &oauth2.Token{
		AccessToken:  creds.Token.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: creds.Token.RefreshToken,
		Expiry:       expiresAt,
	}
	ctx := context.TODO()

	gcfg.OAuth2Config = &oauth2.Config{
		ClientID:     creds.Config.ClientID,
		ClientSecret: creds.Config.ClientSecret,
		Scopes: []string{
			"https://adwords.google.com/api/adwords",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}

	gcfg.Auth = Auth{
		CustomerId:     creds.Auth.CustomerId,
		DeveloperToken: creds.Auth.DeveloperToken,
		Client:         gcfg.OAuth2Config.Client(ctx, gcfg.OAuth2Token),
	}

	return gcfg, nil
}

func NewCredentials(ctx context.Context) (ac AuthConfig, err error) {
	return NewCredentialsFromFile(*configJson, ctx)
}

func NewCredentialsRaw(ctx context.Context, token *oauth2.Token, config *oauth2.Config) (ac AuthConfig) {
	ac.OAuth2Token = token
	ac.OAuth2Config = config
	ac.tokenSource = ac.OAuth2Config.TokenSource(ctx, ac.OAuth2Token)
	ac.Auth.Client = ac.OAuth2Config.Client(ctx, ac.OAuth2Token)
	return ac
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
