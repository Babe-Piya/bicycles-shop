package bicycles

import (
	"bicycles-shop/model"
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

func makeBuyBicycleEndpoint(bs BicyclesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.BuyBicycleRequest)
		response, err = bs.BuyStatusBicycle(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}
