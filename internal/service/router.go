package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/lockton-one/rpc-wrapper-svc/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxRegistryConfig(s.cfg.RegistryConfig()),
			handlers.CtxEthRPCConfig(s.cfg.EthRPCConfig()),
		),
	)
	r.Post("/", handlers.HandleRPC)

	return r
}
