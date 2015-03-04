package gads

import (
	"testing"
)

func testLabelService(t *testing.T) (service *LabelService) {
	return &LabelService{Auth: testAuthSetup(t)}
}

func testLabel(t *testing.T) (Label, func()) {
	s := testLabelService(t)
	labels, err := s.Mutate(
		LabelOperations{
			"ADD": {
				NewTextLabel("Label_" + rand_str(10)),
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	cleanupLabel := func() {
		_, err = s.Mutate(LabelOperations{"REMOVE": labels})
		if err != nil {
			t.Error(err)
		}
	}
	return labels[0], cleanupLabel
}

func TestLabel(t *testing.T) {
	ls := testLabelService(t)
	labels, err := ls.Mutate(
		LabelOperations{
			"ADD": {
				NewTextLabel("Label" + rand_str(10)),
				NewTextLabel("Label" + rand_str(10)),
				NewTextLabel("Label" + rand_str(10)),
				NewTextLabel("Label" + rand_str(10)),
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	foundLabels, _, err := ls.Get(
		Selector{
			Fields: []string{
				"LabelId",
				"LabelName",
				"LabelStatus",
			},
			Predicates: []Predicate{
				{"LabelStatus", "EQUALS", []string{"ENABLED"}},
			},
			Ordering: []OrderBy{
				{"LabelId", "ASCENDING"},
			},
			Paging: &Paging{
				Offset: 0,
				Limit:  500,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("found %d labels\n", len(foundLabels))
	for _, c := range labels {
		func(label Label) {
			for _, foundLabel := range foundLabels {
				if foundLabel.Id == label.Id {
					return
				}
			}
			t.Errorf("label %d not found in \n%#v\n", label.Id, foundLabels)
		}(c)
	}
}
