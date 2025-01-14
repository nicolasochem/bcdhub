package search

import (
	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/models/operation"
)

// Operation -
type Operation struct{}

// GetIndex -
func (o Operation) GetIndex() string {
	return models.DocOperations
}

// GetScores -
func (o Operation) GetScores(search string) []string {
	return []string{
		"entrypoint^8",
		"parameter_strings^7",
		"storage_strings^7",
		"errors.with^6",
		"errors.id^5",
		"source_alias^3",
		"hash",
		"source",
	}
}

// GetFields -
func (o Operation) GetFields() []string {
	return []string{
		"entrypoint",
		"parameter_strings",
		"storage_strings",
		"errors.with",
		"errors.id",
		"source_alias",
		"hash",
		"source",
	}
}

// Parse  -
func (o Operation) Parse(highlight map[string][]string, data []byte) (interface{}, error) {
	var operation operation.Operation
	if err := json.Unmarshal(data, &operation); err != nil {
		return nil, err
	}
	return models.Item{
		Type:       o.GetIndex(),
		Value:      operation.Hash,
		Body:       operation,
		Highlights: highlight,
		Network:    operation.Network,
	}, nil
}
