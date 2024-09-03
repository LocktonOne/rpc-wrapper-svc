package handlers

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/lockton-one/rpc-wrapper-svc/internal/service/requests"
	"gitlab.com/tokene/lockton-one/rpc-wrapper-svc/solidity/generated/allowedcontractregistry"
	"io"
	"net/http"
)

func HandleRPC(w http.ResponseWriter, r *http.Request) {
	request, err := requests.HandleRPC(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if request.Method == "eth_sendRawTransaction" {

		params := []string{}
		if err = json.Unmarshal(request.Params, &params); err != nil {
			Log(r).WithError(errors.Wrap(err, "failed to unmarshal raw message"))
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if len(params[0]) > 2 && params[0][:2] == "0x" {
			params[0] = params[0][2:]
		}

		rawTxBytes, err := hex.DecodeString(params[0])
		if err != nil {
			Log(r).WithError(errors.Wrap(err, "failed to decode raw transaction hex"))
			ape.RenderErr(w, problems.InternalError())
			return
		}

		tx := new(types.Transaction)
		err = rlp.Decode(bytes.NewReader(rawTxBytes), tx)
		if err != nil {
			Log(r).WithError(errors.Wrap(err, "failed to decode raw transaction"))
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if tx.To() == nil {
			var deployData [32]byte
			copy(deployData[:], tx.Data())
			allowedContractRegistry, err := allowedcontractregistry.NewAllowedcontractregistry(RegistryConfig(r).Address, EthRPCConfig(r).EthClient())
			if err != nil {
				Log(r).WithError(err).Debug("can't get contract address")
				ape.RenderErr(w, problems.InternalError())
				return
			}
			ok, err := allowedContractRegistry.IsAllowedToDeploy(nil, deployData)
			if err != nil {
				Log(r).WithError(errors.Wrap(err, "failed to check if contract can be deployed"))
				ape.RenderErr(w, problems.InternalError())
				return
			}
			if !ok {
				Log(r).WithError(errors.Wrap(err, "contract can not be deployed"))
				ape.RenderErr(w, problems.BadRequest(errors.Wrap(err, "contract can not be deployed"))...)
				return
			}
			if _, err = allowedContractRegistry.ToggleDeployedFlag(nil, deployData); err != nil {
				Log(r).WithError(errors.Wrap(err, "failed to toggle deployed flag"))
				ape.RenderErr(w, problems.InternalError())
				return
			}
		}
	}
	reqByte, err := json.Marshal(request)
	if err != nil {
		Log(r).WithError(errors.Wrap(err, "failed to marshal request"))
		ape.RenderErr(w, problems.InternalError())
		return
	}
	res, _ := http.Post(EthRPCConfig(r).Url, "application/json", bytes.NewBuffer(reqByte))
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		Log(r).WithError(errors.Wrap(err, "failed to read rpc response"))
		ape.RenderErr(w, problems.InternalError())
		return
	}
	var unmarshalResp json.RawMessage
	fmt.Println(string(raw))
	if err = json.Unmarshal(raw, &unmarshalResp); err != nil {
		Log(r).WithError(errors.Wrap(err, "failed to unmarshal rpc response"))
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(res.StatusCode)
	ape.Render(w, unmarshalResp)
	return
}
