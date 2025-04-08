package subscan_plugin

import (
	"context"
	"github.com/itering/subscan-plugin/router"
	"github.com/itering/subscan-plugin/storage"
	"github.com/shopspring/decimal"
)

type Plugin interface {
	// Init storage interface
	InitDao(d storage.Dao)

	// Init http router
	InitHttp() []router.Http

	// Receive Extrinsic data when subscribe extrinsic dispatch
	ProcessExtrinsic(*storage.Block, *storage.Extrinsic, []storage.Event) error

	// Receive Extrinsic data when subscribe extrinsic dispatch
	ProcessEvent(*storage.Block, *storage.Event, decimal.Decimal) error

	// Receive Block data when subscribe block
	ProcessBlock(context.Context, *storage.Block) error

	// Mysql tables schema auto migrate
	Migrate()

	// Subscribe Extrinsic with special module
	SubscribeExtrinsic() []string

	// Subscribe Events with special module
	SubscribeEvent() []string

	// Plugins version
	Version() string

	// Set Redis Pool
	SetRedisPool(RedisPool)

	// Plugin enable
	Enable() bool

	// Consumption queue
	ConsumptionQueue() []string

	// ExecWorker exec a task when subscribe queue
	ExecWorker(ctx context.Context, queue, class string, raw interface{}) error
}

type RedisPool interface {
	// HMGet redis hmset
	HMGet(c context.Context, key string, field ...string) (ms map[string]string)
	// GetCacheTtl redis ttl
	GetCacheTtl(ctx context.Context, key string) int
	// HmSetEx redis setex
	HmSetEx(c context.Context, key string, value interface{}, ttl int) (err error)
}
