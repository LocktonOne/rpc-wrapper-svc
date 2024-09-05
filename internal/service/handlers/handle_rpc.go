package handlers

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/lockton-one/rpc-wrapper-svc/internal/service/requests"
	"gitlab.com/tokene/lockton-one/rpc-wrapper-svc/solidity/generated/allowedcontractregistry"
	"golang.org/x/crypto/sha3"
	"io"
	"net/http"
)

func HandleRPC(w http.ResponseWriter, r *http.Request) {
	// Read the body into a byte slice
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		Log(r).WithError(err).Info("failed to get body bytes")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// Restore the body so it can be read again
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	request, err := requests.HandleRPC(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if request.Method == "eth_sendRawTransaction" {
		Log(r).Debug("THIS IS eth_sendRawTransaction")
		spew.Dump(request.Params)

		params := []string{}
		if err = json.Unmarshal(request.Params, &params); err != nil {
			Log(r).WithError(err).Info("failed to unmarshal raw message")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if len(params[0]) > 2 && params[0][:2] == "0x" {
			params[0] = params[0][2:]
		}
		Log(r).Debug(params[0])
		rawTxBytes, err := hex.DecodeString(params[0])
		if err != nil {
			Log(r).WithError(err).Info("failed to decode raw transaction hex")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		tx := new(types.Transaction)

		if len(rawTxBytes) > 0 && rawTxBytes[0] >= 0x01 && rawTxBytes[0] <= 0x7f {
			err = tx.UnmarshalBinary(rawTxBytes)
			if err != nil {
				Log(r).WithError(err).Info("failed to unmarshal typed transaction")
				ape.RenderErr(w, problems.InternalError())
				return
			}
		} else {
			if err = rlp.Decode(bytes.NewReader(rawTxBytes), tx); err != nil {
				Log(r).WithError(err).Info("failed to decode raw transaction")
				ape.RenderErr(w, problems.InternalError())
				return
			}
		}

		Log(r).Debug(tx.Hash())
		if tx.To() == nil {
			Log(r).Debug("THIS IS deploy")
			hash := sha3.NewLegacyKeccak256()
			hash.Write(tx.Data())
			hashedData := hash.Sum(nil)
			var deployData [32]byte
			copy(deployData[:], hashedData)
			Log(r).Debug("hashedData", hashedData)
			allowedContractRegistry, err := allowedcontractregistry.NewAllowedcontractregistry(RegistryConfig(r).Address, EthRPCConfig(r).EthClient())
			if err != nil {
				Log(r).WithError(err).Debug("can't get contract address")
				ape.RenderErr(w, problems.InternalError())
				return
			}
			ok, err := allowedContractRegistry.IsAllowedToDeploy(nil, deployData)
			if err != nil {
				Log(r).WithError(err).Info("failed to check if contract can be deployed")
				ape.RenderErr(w, problems.InternalError())
				return
			}
			if !ok {
				Log(r).Info("contract can not be deployed")
				ape.RenderErr(w, problems.BadRequest(errors.Wrap(err, "contract can not be deployed"))...)
				return
			}
			//if _, err = allowedContractRegistry.ToggleDeployedFlag(nil, deployData); err != nil {
			//	Log(r).WithError(err).Info("failed to toggle deployed flag")
			//	ape.RenderErr(w, problems.InternalError())
			//	return
			//}
		}
	}
	//reqByte, err := json.Marshal(request)
	//if err != nil {
	//	Log(r).WithError(err).Info("failed to marshal request"))
	//	ape.RenderErr(w, problems.InternalError())
	//	return
	//}
	res, err := http.Post(EthRPCConfig(r).Url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		Log(r).WithError(err).Info("failed to send rpc transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		Log(r).WithError(err).Info("failed to read rpc response")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	var unmarshalResp json.RawMessage
	Log(r).Debug(string(raw))
	if err = json.Unmarshal(raw, &unmarshalResp); err != nil {
		Log(r).WithError(err).Info("failed to unmarshal rpc response")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(res.StatusCode)
	ape.Render(w, unmarshalResp)
	return
}
