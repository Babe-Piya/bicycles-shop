package bicycles

import (
	"bicycles-shop/model"
	"bicycles-shop/repository"
	"context"
)

type BicyclesService interface {
	GetListBicycles(ctx context.Context) (model.BicyclesResponse, error)
	BuyStatusBicycle(ctx context.Context, req model.BuyBicycleRequest) (model.BuyBicycleResponse, error)
	CreateBicycle(req model.BicycleRequest) (model.CreateBicycleResponse, error)
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

func (bs *bicyclesService) BuyStatusBicycle(ctx context.Context, req model.BuyBicycleRequest) (model.BuyBicycleResponse, error) {
	var response model.BuyBicycleResponse

	buyBicycle, err := bs.repo.GetBicycleCanBuy(ctx, req.BicycleID)
	if err != nil {
		response = model.BuyBicycleResponse{
			Success: false,
			Message: err.Error(),
		}
		return response, err
	}
	bicycleData := model.BicycleData{
		ID:    buyBicycle.ID,
		Brand: buyBicycle.Brand,
		Model: buyBicycle.Model,
		Price: buyBicycle.Price,
	}

	err = bs.repo.UpdateBuyStatusBicycle(ctx, req.BicycleID)
	if err != nil {
		response = model.BuyBicycleResponse{
			Success: false,
			Message: err.Error(),
		}
		return response, err
	}

	createBuyer := model.Buyers{
		BicycleID: req.BicycleID,
		Name:      req.Name,
		Address:   req.Address,
		Tel:       req.Tel,
	}
	err = bs.repo.InsertBuyer(ctx, createBuyer)
	if err != nil {
		response = model.BuyBicycleResponse{
			Success: false,
			Message: err.Error(),
		}
		return response, err
	}

	response = model.BuyBicycleResponse{
		Success: true,
		Data:    bicycleData,
	}
	return response, nil
}

func (bs *bicyclesService) CreateBicycle(req model.BicycleRequest) (model.CreateBicycleResponse, error) {
	bicycleReq := model.Bicycles{
		Brand:       req.Brand,
		Model:       req.Model,
		Description: req.Description,
		Price:       req.Price,
	}

	bicycleRes, err := bs.repo.InsertBicycle(bicycleReq)
	if err != nil {
		return model.CreateBicycleResponse{}, nil
	}
	res := model.CreateBicycleResponse{
		Success: true,
		Data:    bicycleRes,
	}

	return res, nil
}
