package service

import (
	"context"
	"fmt"
	"github.com/gookit/goutil/errorx"
	"microservice-test/proto/price"
	"time"
)

const specialDay = 20

type Service struct {
	price.UnimplementedPriceServiceServer
}

func New() *Service {
	return &Service{}
}

func (ps *Service) GetPrice(ctx context.Context, in *price.PriceRequest) (*price.PriceResponse, error) {
	fmt.Println("[RPC] GetPrice Called With Request ", in)
	date := in.Date
	if len(date) == 0 {
		return nil, errorx.Wrap(errorx.New("Date can't be empty"), fmt.Sprintf("[RPC] GetPrice ValidateAll Failed %v", date))
	}

	timeParsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, errorx.Wrap(errorx.New("Date must be in format 2006-01-02"), fmt.Sprintf("[RPC] GetPrice Parse Time Failed %v", err))
	}

	var priceForTheDay int64 = 100000
	if timeParsed.Day() == specialDay {
		priceForTheDay = 200000
	}

	res := &price.PriceResponse{
		Date:  date,
		Price: priceForTheDay,
	}
	fmt.Println("[RPC] GetPrice Response ", res)

	return res, nil
}
