// Package gads provides a wrapper for the Google Adwords SOAP API.
// In order to access the API you will need to sign up for an MMC
// account[1], get a developer token[2], setup authentication[3].  The is
// a tool in the setup_oauth2 directory that will setup a configuration
// file.
//
// The package is comprised of services used to manipulate various
// adwords structures.  To access a service you need to create an
// gads.Auth and parse it to the service initializer, then can call
// the service methods on the service object.
//
//     authConf, err := NewCredentials(context.TODO())
//     campaignService := gads.NewCampaignService(&authConf.Auth)
//
//     campaigns, totalCount, err := cs.Get(
//       gads.Selector{
//         Fields: []string{
//           "Id",
//           "Name",
//           "Status",
//         },
//       },
//     )
//
// 1. http://www.google.com/adwords/myclientcenter/
//
// 2. https://developers.google.com/adwords/api/docs/signingup
//
// 3. https://developers.google.com/adwords/api/docs/guides/authentication
package gads
