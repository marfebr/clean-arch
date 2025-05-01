package usecase

import (
	"github.com/marfebr/cleanarch/internal/entity"
	"github.com/marfebr/cleanarch/pkg/events"
)

type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}
type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreate     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreate events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {

	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreate:     OrderCreate,
		EventDispatcher: EventDispatcher,
	}

}

func (c *CreateOrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	order.CalculateFinalPrice()
	if err := c.OrderRepository.Save(&order); err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
	c.OrderCreate.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderCreate)

	return dto, nil
}
