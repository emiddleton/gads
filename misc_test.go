package gads

import (
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

func ExampleMediaService_Upload() {
	// load image into []byte
	imageUrl := "http://www.google.com/intl/en/adwords/select/images/samples/inline.jpg"
	resp, err := http.Get(imageUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// load credentials from
	authConf, _ := NewCredentials(oauth2.NoContext)
	ms := NewMediaService(&authConf.Auth)

	images, err := ms.Upload([]Media{NewImage("image1", "IMAGE", "IMAGE_JPEG", body)})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", images)
}
