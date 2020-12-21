package config

import (
	"github.com/baking-bad/bcdhub/internal/aws"
	"github.com/baking-bad/bcdhub/internal/contractparser/kinds"
	"github.com/baking-bad/bcdhub/internal/database"
	"github.com/baking-bad/bcdhub/internal/elastic"
	"github.com/baking-bad/bcdhub/internal/mq"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/baking-bad/bcdhub/internal/pinata"
	"github.com/baking-bad/bcdhub/internal/tzkt"
	"github.com/pkg/errors"
)

// Context -
type Context struct {
	DB           database.DB
	ES           elastic.IElastic
	MQ           mq.Mediator
	AWS          *aws.Client
	RPC          map[string]noderpc.INode
	TzKTServices map[string]tzkt.Service
	Pinata       pinata.Service

	Config     Config
	SharePath  string
	TzipSchema string

	Interfaces map[string]kinds.ContractKind
	Domains    map[string]string
}

// NewContext -
func NewContext(opts ...ContextOption) *Context {
	ctx := &Context{}

	for _, opt := range opts {
		opt(ctx)
	}
	return ctx
}

// GetRPC -
func (ctx *Context) GetRPC(network string) (noderpc.INode, error) {
	if rpc, ok := ctx.RPC[network]; ok {
		return rpc, nil
	}
	return nil, errors.Errorf("Unknown rpc network %s", network)
}

// GetTzKTService -
func (ctx *Context) GetTzKTService(network string) (tzkt.Service, error) {
	if rpc, ok := ctx.TzKTServices[network]; ok {
		return rpc, nil
	}
	return nil, errors.Errorf("Unknown tzkt service network %s", network)
}

// Close -
func (ctx *Context) Close() {
	if ctx.MQ != nil {
		ctx.MQ.Close()
	}
	if ctx.DB != nil {
		ctx.DB.Close()
	}
}
