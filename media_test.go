package gads

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func testMediaService(t *testing.T) (service *MediaService) {
	return &MediaService{Auth: testAuthSetup(t)}
}

func TestMedia(t *testing.T) {

	// load image into []byte
	imageUrl := "http://www.google.com/intl/en/adwords/select/images/samples/inline.jpg"
	resp, err := http.Get(imageUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	ms := testMediaService(t)
	images, err := ms.Upload(
		[]Media{
			NewImage("image1", "IMAGE", "IMAGE_JPEG", body),
			NewImage("image2", "IMAGE", "IMAGE_JPEG", body),
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", images)

	var pageSize int64 = 500
	var offset int64 = 0
	paging := Paging{
		Offset: offset,
		Limit:  pageSize,
	}
	for {
		medias, totalCount, err := ms.Get(
			Selector{
				Fields: []string{
					"MediaId",
					"Height",
					"Width",
					"MimeType",
					"Urls",
				},
				Predicates: []Predicate{
					{"Type", "IN", []string{"IMAGE", "VIDEO"}},
				},
				Paging: &paging,
			},
		)
		if err != nil {
			fmt.Printf("Error occured finding medias")
		}
		for _, m := range medias {
			for _, d := range m.Dimensions {
				if d.Name == "FULL" {
					fmt.Printf("Entry ID %d with dimensions %dx%d and MIME type is '%s'\n", m.Id, d.Height, d.Width, m.MimeType)
				}
			}
		}
		// Increment values to request the next page.
		offset += pageSize
		paging.Offset = offset
		if totalCount < offset {
			fmt.Printf("\tFound %d entries.", totalCount)
			break
		}
	}

}
