package models

// Histogram filter kinds
const (
	HistogramFilterKindExists     = "exists"
	HistogramFilterKindMatch      = "match"
	HistogramFilterKindIn         = "in"
	HistogramFilterKindAddresses  = "address"
	HistogramFilterDexEnrtypoints = "dex_entrypoints"
)

// HistogramContext -
type HistogramContext struct {
	Indices  []string
	Period   string
	Function struct {
		Name  string
		Field string
	}
	Filters []HistogramFilter
}

// HasFunction -
func (ctx HistogramContext) HasFunction() bool {
	return ctx.Function.Name != "" && ctx.Function.Field != ""
}

// HistogramFilter -
type HistogramFilter struct {
	Field string
	Value interface{}
	Kind  string
}

// HistogramOption -
type HistogramOption func(*HistogramContext)

// WithHistogramIndices -
func WithHistogramIndices(indices ...string) HistogramOption {
	return func(h *HistogramContext) {
		h.Indices = indices
	}
}

// WithHistogramFunction -
func WithHistogramFunction(function, field string) HistogramOption {
	return func(h *HistogramContext) {
		h.Function = struct {
			Name  string
			Field string
		}{function, field}
	}
}

// WithHistogramFilters -
func WithHistogramFilters(filters []HistogramFilter) HistogramOption {
	return func(h *HistogramContext) {
		h.Filters = filters
	}
}
