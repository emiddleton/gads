package gads

import (
	"encoding/xml"
	"fmt"
)

func findAttr(as []xml.Attr, name xml.Name) (value string, err error) {
	for _, a := range as {
		an := a.Name
		if an.Space == name.Space && an.Local == name.Local {
			return a.Value, nil
		}
	}
	return "", fmt.Errorf("attribute %#v not found", name)
}
