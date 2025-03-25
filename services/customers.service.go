package services

import (
	"context"
	"strconv"

	"github.com/milkyway/gin_beginer/models"
	"github.com/milkyway/gin_beginer/repositories"
)

type CustomerService struct {
	CustomersRepo repositories.CustomersRepo
}

func (cs CustomerService) GetInfoByRegistry(ctx context.Context, cust_registry string) (models.Customers, error) {
	var result models.Customers
	var err error = nil
	result, err = cs.CustomersRepo.FetchByRegistry(ctx, cust_registry)

	// note: handle error di sisi service
	if err != nil {
		return result, err
	}

	return result, err
}

// NOTE: parsing array request from postman
// NOTE: must be transaction!
func (cs CustomerService) BuyBooks(ctx context.Context, books_repo repositories.BooksRepo, order_req []OrderBookRequest) ([]models.Books, error) {

	var order_validation []models.Books

	for _, book_info := range order_req {

		id, err := strconv.Atoi(book_info.ID)
		if err != nil {
			return order_validation, err
		}

		result_fetch, err := books_repo.FetchBookByID(ctx, id)
		if err != nil {
			return order_validation, err
		}

		order_validation = append(order_validation, result_fetch)
	}
	// harusnya doing transaction here!
	return order_validation, nil
}

func (cs CustomerService) SetDiscount(float64) float64 {
	return 0.25
}
