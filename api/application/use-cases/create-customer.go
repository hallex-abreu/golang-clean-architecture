package usecases

import (
	"errors"

	"github.com/hallex-abreu/golang-clean-architecture/api/application/repositories"
	"github.com/hallex-abreu/golang-clean-architecture/api/domain"
)

type CreateCustomerRequest struct {
	Name  string
	Email string
	Phone string
}

func CreateCustomer(customerRequest CreateCustomerRequest, customerRepository repositories.CustomerRepository) (domain.Customer, error) {
	exist_customer := customerRepository.FindCustomerByEmail(customerRequest.Email)

	if exist_customer != nil {
		return domain.Customer{}, errors.New("Already customer with this email")
	}

	customer := domain.Customer{
		Name:  customerRequest.Name,
		Email: customerRequest.Email,
		Phone: customerRequest.Phone,
	}

	customerRepository.CreateCustomer(customer)

	return customer, nil
}
