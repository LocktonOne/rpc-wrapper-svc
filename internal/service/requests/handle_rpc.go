package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/lockton-one/rpc-wrapper-svc/resources"
	"net/http"
)

func HandleRPC(r *http.Request) (resources.RpcRequest, error) {
	var request resources.RpcRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	return request, nil
}
