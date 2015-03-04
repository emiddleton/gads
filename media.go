package gads

import (
	"encoding/base64"
	"encoding/xml"
)

type MediaService struct {
	Auth
}

func NewMediaService(auth *Auth) *MediaService {
	return &MediaService{Auth: *auth}
}

type Dimensions struct {
	Name   string `xml:"key"` // "FULL", "SHRUNKEN", "PREVIEW", "VIDEO_THUMBNAIL"
	Width  int64  `xml:"value>width"`
	Height int64  `xml:"value>height"`
}

type ImageUrl struct {
}

// Media represents an audio, image or video file.
type Media struct {
	Type         string       `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Id           int64        `xml:"mediaId,omitempty"`
	MediaType    string       `xml:"type"` // "AUDIO", "DYNAMIC_IMAGE", "ICON", "IMAGE", "STANDARD_ICON", "VIDEO"
	ReferenceId  int64        `xml:"referenceId,omitempty"`
	Dimensions   []Dimensions `xml:"dimensions"`
	Urls         []ImageUrl   `xml:"urls"`
	MimeType     string       `xml:"mimeType,omitempty"` // "IMAGE_JPEG", "IMAGE_GIF", "IMAGE_PNG", "FLASH", "TEXT_HTML", "PDF", "MSWORD", "MSEXCEL", "RTF", "AUDIO_WAV", "AUDIO_MP3"
	SourceUrl    string       `xml:"sourceUrl,omitempty"`
	Name         string       `xml:"name"`
	FileSize     int64        `xml:"fileSize,omitempty"`   // File size in bytes
	CreationTime string       `xml:"createTime,omitempty"` // format is YYYY-MM-DD HH:MM:SS+TZ

	// Audio / Video
	DurationMillis   int64  `xml:"durationMillis,omitempty"`
	StreamingUrl     string `xml:"streamingUrl,omitempty"`
	ReadyToPlayOnWeb bool   `xml:"readyToPlayOnTheWeb,omitempty"`

	// Image
	Data string `xml:"data,omitempty"` // base64Binary encoded raw image data

	// Video
	IndustryStandardCommercialIdentifier string `xml:"industryStandardCommercialIdentifier,omitempty"`
	AdvertisingId                        string `xml:"advertisingId,omitempty"`
	YouTubeVideoIdString                 string `xml:"youTubeVideoIdString,omitempty"`
}

func NewAudio(name, mediaType, mimeType string) (image Media) {
	return Media{
		Name: name,
		Type: "Audio",
	}
}

func NewImage(name, mediaType, mimeType string, data []byte) (image Media) {
	imageData := base64.StdEncoding.EncodeToString(data)
	return Media{
		Name:      name,
		Type:      "Image",
		MediaType: mediaType,
		Data:      imageData,
	}
}

func NewVideo(mediaType string) (image Media) {
	return Media{
		Type: "Video",
	}
}

func (s *MediaService) Get(selector Selector) (medias []Media, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		mediaServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "get",
			},
			Sel: selector,
		},
	)
	if err != nil {
		return medias, totalCount, err
	}
	getResp := struct {
		Size   int64   `xml:"rval>totalNumEntries"`
		Medias []Media `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return medias, totalCount, err
	}
	return getResp.Medias, getResp.Size, err
}

func (s *MediaService) Query(query string) (medias []Media, totalCount int64, err error) {
	return medias, totalCount, ERROR_NOT_YET_IMPLEMENTED
}

func (s *MediaService) Upload(medias []Media) (uploadedMedias []Media, err error) {
	upload := struct {
		XMLName xml.Name
		Medias  []Media `xml:"media"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "upload",
		},
		Medias: medias,
	}
	respBody, err := s.Auth.request(mediaServiceUrl, "upload", upload)
	if err != nil {
		return uploadedMedias, err
	}
	uploadResp := struct {
		Medias []Media `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &uploadResp)
	if err != nil {
		return uploadedMedias, err
	}
	return uploadResp.Medias, err
}
