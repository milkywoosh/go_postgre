package services

import (
	"context"

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
