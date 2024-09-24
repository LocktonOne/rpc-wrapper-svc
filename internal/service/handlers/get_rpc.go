package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetRPC(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w, r)
	res, err := http.Get(EthRPCConfig(r).Url)
	if err != nil {
		Log(r).WithError(err).Info("failed to send get rpc method")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(res.StatusCode)
	return
}
