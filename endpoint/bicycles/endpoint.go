package bicycles

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeGetListBicyclesEndpoint(bs BicyclesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		response, err = bs.GetListBicycles(ctx)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}
