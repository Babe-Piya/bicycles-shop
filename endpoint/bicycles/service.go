package bicycles

import (
	"bicycles-shop/model"
	"bicycles-shop/repository"
	"context"
)

type BicyclesService interface {
	GetListBicycles(ctx context.Context) (model.BicyclesResponse, error)
}

type bicyclesService struct {
	repo repository.BicyclesRepository
}

func NewService(repo repository.BicyclesRepository) BicyclesService {
	return &bicyclesService{
		repo: repo,
	}
}

func (bs *bicyclesService) GetListBicycles(ctx context.Context) (model.BicyclesResponse, error) {
	var response model.BicyclesResponse
	//log := zap.S().With("trace_id", data.CorrelationId, "event", "SendWelcomeEmail")

	list, err := bs.repo.GetListBicycles(ctx)
	if err != nil {
		response = model.BicyclesResponse{
			Success: false,
			Message: err.Error(),
		}
	}
	response = model.BicyclesResponse{
		Success: true,
		Data:    list,
	}
	return response, nil
}
