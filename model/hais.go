package model

import (
	"github.com/fatih/structs"
)

type Hais struct {
	ID      int    `json:"id" structs:"id" bson:"_id" db:"h_id"`
	City    string `json:"city" structs:"city" bson:"city" db:"city"`
	HaiName string `json:"haiName" structs:"haiName" bson:"haiName" db:"hai"`
}

// Map converts structs to a map representation
func (h *Hais) Map() map[string]interface{} {
	return structs.Map(h)
}

// Names returns the field names of Subscription model.
func (h *Hais) Names() []string {
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
