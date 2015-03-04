// setup_oauth2 is a tool for creating a gads configuration file config.json from
// the installed application credential stored in credentials.json.  The utility will
// open the Google consent page asking you to grant permission to the application.  Login
// as your MCC account user. Once you have granted permission you will be provided with an
// authorization code to copy back to the tool.
//
// The generated config.json will look something like this
//
//     {
//         "oauth2.Config": {
//             "ClientID": "4585432543254323-f4qfewtg2qtg5esy24t45h.apps.googleusercontent.com",
//             "ClientSecret": "fa74ehgyjhtrrjtbrsu56hHjhhrtger",
//             "Endpoint": {
//                 "AuthURL": "https://accounts.google.com/o/oauth2/auth",
//                 "TokenURL": "https://accounts.google.com/o/oauth2/token"
//             },
//             "RedirectURL": "oob",
//             "Scopes": [
//                 "https://adwords.google.com/api/adwords"
//             ]
//         },
//         "oauth2.Token": {
//             "access_token": "jfdsalkfjdskalfjdaksfdasfdsahrtsrgf",
//             "token_type": "Bearer",
//             "refresh_token": "g65wurefej87ruy4fcyfdsafdsafdsafsdaf4fu",
//             "expiry": "2015-03-05T00:13:23.382907238+09:00"
//         },
//         "gads.Auth": {
//             "CustomerId": "INSERT_YOUR_CLIENT_CUSTOMER_ID_HERE",
//             "DeveloperToken": "INSERT_YOUR_DEVELOPER_TOKEN_HERE",
//             "UserAgent": "tests (Golang 1.4 github.com/emiddleton/gads)"
//         }
//     }
//
// You will need to add you customer id (which you can find by logging into the
// "My Client Center" http://www.google.com/adwords/myclientcenter/) and your
// developer token.
//
// You can find details on how to create the credentials here.
// https://developers.google.com/adwords/api/docs/guides/authentication
// You can download credentials.json from the google developer console under
//
//     "API's & auth" > "Credentials" > "OAuth" > "Client ID for native application" > "Download JSON"
//
package main
