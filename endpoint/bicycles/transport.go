package bicycles

import (
	"bicycles-shop/model"
	"context"
	"encoding/json"
	"errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func MakeGetListBicyclesHandler(bs BicyclesService) http.Handler {
	return kithttp.NewServer(
		makeGetListBicyclesEndpoint(bs),
		decodeGetListBicyclesRequest,
		encodeResponse,
	)
}

func decodeGetListBicyclesRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	if _, ok := response.(error); ok {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func MakeBuyBicycleHandler(bs BicyclesService) http.Handler {
	return kithttp.NewServer(
		makeBuyBicycleEndpoint(bs),
		decodeBuyBicycleRequest,
		encodeResponse,
	)
}

func decodeBuyBicycleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BuyBicycleRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func MakeCreateBicycleHandler(bs BicyclesService) http.Handler {
	return kithttp.NewServer(
		makeCreateBicycleEndpoint(bs),
		decodeCreateBicycleRequest,
		encodeResponse,
	)
}

func decodeCreateBicycleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BicycleRequest
	auth := r.Header.Get("Authorization")
	if auth != "tokensomething" {
		return nil, errors.New("authorization failed")
	}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}
