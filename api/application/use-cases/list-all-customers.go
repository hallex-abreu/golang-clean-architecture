package usecases

import (
	"errors"
	"math"

	"github.com/hallex-abreu/golang-clean-architecture/api/application/repositories"
	"github.com/hallex-abreu/golang-clean-architecture/api/domain"
)

type ListAllCustomersRequest struct {
	Page   int
	Limit  int
	Filter string
}

type ListAllCustomersResponse struct {
	Content        []domain.Customer
	Page           int
	Limit          int
	TotalElements  int
	NumberElements int
	TotalPages     float64
}

func ListAllCustomers(listAllCustomersRequest ListAllCustomersRequest, customerRepository repositories.CustomerRepository) (ListAllCustomersResponse, error) {
	customers := customerRepository.FindAllCustomers(listAllCustomersRequest.Page, listAllCustomersRequest.Limit, listAllCustomersRequest.Filter)

	if len(*customers) == 0 {
		return ListAllCustomersResponse{}, errors.New("Does not exists customers")
	}

	count := customerRepository.GetCountCustomers(listAllCustomersRequest.Filter)

	var total_elements = math.Ceil(float64(count) / float64(listAllCustomersRequest.Limit))

	return ListAllCustomersResponse{
		Content:        *customers,
		Page:           listAllCustomersRequest.Page,
		Limit:          listAllCustomersRequest.Limit,
		TotalElements:  count,
		NumberElements: len(*customers),
		TotalPages:     total_elements,
	}, nil
}
