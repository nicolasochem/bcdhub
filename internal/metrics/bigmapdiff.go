package metrics

import (
	"encoding/json"

	"github.com/baking-bad/bcdhub/internal/contractparser/stringer"
	"github.com/baking-bad/bcdhub/internal/models"
)

// SetBigMapDiffsStrings -
func (h *Handler) SetBigMapDiffsStrings(bmd *models.BigMapDiff) error {
	keyBytes, err := json.Marshal(bmd.Key)
	if err != nil {
		return err
	}
	bmd.KeyStrings = stringer.Get(string(keyBytes))
	bmd.ValueStrings = stringer.Get(bmd.Value)
	return nil
}
