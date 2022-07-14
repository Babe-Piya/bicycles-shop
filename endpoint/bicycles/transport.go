package bicycles

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func MakeGetSubscriptionPackageHandler(bs BicyclesService) http.Handler {
	return kithttp.NewServer(
		makeGetListBicyclesEndpoint(bs),
		decodeGetSubscriptionPackageRequest,
		encodeGetSubscriptionPackageResponse,
	)
}

func decodeGetSubscriptionPackageRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeGetSubscriptionPackageResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	if _, ok := response.(error); ok {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
