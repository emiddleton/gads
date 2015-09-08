# gads

Package gads provides a wrapper for the Google Adwords SOAP API.

## installation

~~~
go get github.com/emiddleton/gads
~~~

## setup

In order to access the API you will need to sign up for an MMC
account[1], get a developer token[2] and setup authentication[3].
There is a tool in the setup_oauth2 directory that will help you
setup a configuration file.

1. http://www.google.com/adwords/myclientcenter/
2. https://developers.google.com/adwords/api/docs/signingup
3. https://developers.google.com/adwords/api/docs/guides/authentication

## usage

The package is comprised of services used to manipulate various
adwords structures.  To access a service you need to create an
gads.Auth and parse it to the service initializer, then can call
the service methods on the service object.

~~~ go
     authConf, err := NewCredentials(context.TODO())
     campaignService := gads.NewCampaignService(&authConf.Auth)

     campaigns, totalCount, err := campaignService.Get(
       gads.Selector{
         Fields: []string{
           "Id",
           "Name",
           "Status",
         },
       },
     )
~~~

> Note: This package is a work-in-progress, and may occasionally
> make backwards-incompatible changes.

See godoc for further documentation and examples.

* [godoc.org/github.com/emiddleton/gads](https://godoc.org/github.com/emiddleton/gads)

## about

Gads is developed by [Edward Middleton](https://blog.vortorus.net/)
