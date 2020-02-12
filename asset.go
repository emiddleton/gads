package gads

import (
	"encoding/base64"
	"encoding/xml"
)

type AssetService struct {
	Auth
}

func NewAssetService(auth *Auth) *AssetService {
	return &AssetService{Auth: *auth}
}

type ImageDimensionInfo struct {
	Url    string  `xml:"imageUrl"` // "FULL", "SHRUNKEN", "PREVIEW", "VIDEO_THUMBNAIL"
	Width  float64 `xml:"imageHeight"`
	Height float64 `xml:"imageWidth"`
}

// Media represents an audio, image or video file.
type Asset struct {
	Id        int64  `xml:"assetId,omitempty"`
	Name      string `xml:"assetName,omitempty"`
	Type      string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	AssetType string `xml:"assetSubtype"` // "TEXT", "IMAGE", "YOUTUBE_VIDEO", "MEDIA_BUNDLE"
	Status    string `xml:"assetStatus,omitempty"`

	// Media Bundle
	MediaData string `xml:"mediaBundleData,omitempty"` // base64Binary encoded raw media bundle data

	// Image
	ImgData      string             `xml:"imageData,omitempty"` // base64Binary encoded raw image data
	ImgSize      float64            `xml:"imageFileSize,omitempty"`
	MimeType     string             `xml:"imageMimeType,omitempty"` // "IMAGE_JPEG", "IMAGE_GIF", "IMAGE_PNG"
	FullSizeInfo ImageDimensionInfo `xml:"fullSizeInfo,omitempty"`

	// Youtube Video
	YoutubeVideoId string `xml:"youTubeVideoId,omitempty"` // This is the 11 char string value used in the Youtube video URL
}

func NewMediaBundle(name, assetType string, data []byte) (image Asset) {
	mediaData := base64.StdEncoding.EncodeToString(data)
	return Asset{
		Name:      name,
		Type:      "media_bundle",
		AssetType: assetType,
		MediaData: mediaData,
	}
}

func NewImageAsset(name, assetType, mimeType string, data []byte) (image Asset) {
	imageData := base64.StdEncoding.EncodeToString(data)
	return Asset{
		Name:      name,
		Type:      "Image",
		AssetType: assetType,
		ImgData:   imageData,
		MimeType:  mimeType,
	}
}

func NewVideoAsset(assetType string) (image Asset) {
	return Asset{
		Type:      "Youtube_Video",
		AssetType: assetType,
	}
}

func (s *AssetService) Get(selector Selector) (assets []Asset, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "selector"}
	respBody, err := s.Auth.request(
		assetServiceUrl,
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
		return assets, totalCount, err
	}
	getResp := struct {
		Size   int64   `xml:"rval>totalNumEntries"`
		Assets []Asset `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return assets, totalCount, err
	}
	return getResp.Assets, getResp.Size, err
}
