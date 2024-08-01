package handlers

import (
	"context"
	"gitlab.com/tokene/lockton-one/rpc-wrapper-svc/internal/config"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	registryConfigCtxKey
	ethrpcConfigCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxRegistryConfig(entry *config.RegistryConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, registryConfigCtxKey, entry)
	}
}
func RegistryConfig(r *http.Request) *config.RegistryConfig {
	return r.Context().Value(registryConfigCtxKey).(*config.RegistryConfig)
}

func CtxEthRPCConfig(entry *config.EthRPCConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ethrpcConfigCtxKey, entry)
	}
}

func EthRPCConfig(r *http.Request) *config.EthRPCConfig {
	return r.Context().Value(ethrpcConfigCtxKey).(*config.EthRPCConfig)
}
