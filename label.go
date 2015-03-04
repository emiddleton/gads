package gads

import (
	"encoding/xml"
)

type LabelService struct {
	Auth
}

func NewLabelService(auth *Auth) *LabelService {
	return &LabelService{Auth: *auth}
}

// Label represents a label.
type Label struct {
	Type   string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Id     int64  `xml:"id,omitempty"`
	Name   string `xml:"name"`
	Status string `xml:"status,omitempty"`
}

// NewTextLabel returns an new Label struct for creating a new TextLabel.
func NewTextLabel(name string) Label {
	return Label{
		Type: "TextLabel",
		Name: name,
	}
}

// LabelOperations is a map of operations to perform on Label's
type LabelOperations map[string][]Label

// Get returns an array of Label's and the total number of Label's matching
// the selector.
//
// Example
//
//   labels, totalCount, err := labelService.Get(
//     Selector{
//       Fields: []string{"LabelId","LabelName","LabelStatus"},
//       Predicates: []Predicate{
//         {"LabelStatus", "EQUALS", []string{"ENABLED"}},
//       },
//     },
//   )
//
// Selectable fields are
//   "LabelId", "LabelName", "LabelStatus"
//
// filterable fields are
//   "LabelId", "LabelName", "LabelStatus"
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/LabelService#get
//
func (s LabelService) Get(selector Selector) (labels []Label, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		labelServiceUrl,
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
		return labels, totalCount, err
	}
	getResp := struct {
		Size   int64   `xml:"rval>totalNumEntries"`
		Labels []Label `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return labels, totalCount, err
	}
	return getResp.Labels, getResp.Size, err
}

// Mutate allows you to add, modify and remove labels, returning the
// modified labels.
//
// Example
//
//  labels, err := labelService.Mutate(
//    LabelOperations{
//      "ADD": {
//        NewTextLabel("Label1"),
//        NewTextLabel("Label2"),
//      },
//      "SET": {
//        modifiedLabel,
//      },
//      "REMOVE": {
//        Label{Type:"TextLabel",Id:10},
//      },
//    },
//  )
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/LabelService#mutate
//
func (s *LabelService) Mutate(labelOperations LabelOperations) (labels []Label, err error) {
	type labelOperation struct {
		Action string `xml:"operator"`
		Label  Label  `xml:"operand"`
	}
	operations := []labelOperation{}
	for action, labels := range labelOperations {
		for _, label := range labels {
			operations = append(operations,
				labelOperation{
					Action: action,
					Label:  label,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []labelOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: operations,
	}
	respBody, err := s.Auth.request(labelServiceUrl, "mutate", mutation)
	if err != nil {
		return labels, err
	}
	mutateResp := struct {
		Labels []Label `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return labels, err
	}
	return mutateResp.Labels, err
}

// Query is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/LabelService#query
//
func (s *LabelService) Query(query string) (labels []Label, totalCount int64, err error) {
	return labels, totalCount, ERROR_NOT_YET_IMPLEMENTED
}
