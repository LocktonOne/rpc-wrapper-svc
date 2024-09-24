package service

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "DNT", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Range", "Signature", "Account-Id"},
			ExposedHeaders:   []string{"Content-Length", "Content-Range"},
			AllowCredentials: true,
		}),
	)
	r.Post("/", handlers.HandleRPC)
	return r
}
