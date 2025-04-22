package usecase

import (
	"log/slog"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrderUserCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUserCase(orderRepo entity.OrderRepositoryInterface) *ListOrderUserCase {

	return &ListOrderUserCase{
		OrderRepository: orderRepo,
	}
}
func (l *ListOrderUserCase) Execute() ([]OrderOutputDTO, error) {

	orders, err := l.OrderRepository.List()
	slog.Debug("List orders Repo", "orders", orders)
	if err != nil {
		return nil, err
	}
	entityOrders := make([]OrderOutputDTO, len(orders))
	for i, order := range orders {
		entityOrders[i] = OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}
	return entityOrders, nil
}
