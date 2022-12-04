package model

import (
	"github.com/fatih/structs"
)

type HaiCounter struct {
	City       string `json:"city" structs:"city" bson:"city" db:"city"`
	Haicounter int    `json:"haicounter" structs:"haicounter" bson:"haicounter" db:"haicounter"`
}

// Map converts structs to a map representation
func (h *HaiCounter) Map() map[string]interface{} {
	return structs.Map(h)
}

// Names returns the field names of Subscription model.
func (h *HaiCounter) Names() []string {
	fields := structs.Fields(h)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
