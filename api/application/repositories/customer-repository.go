package repositories

import "github.com/hallex-abreu/golang-clean-architecture/api/domain"

type CustomerRepository interface {
	FindCustomerByEmail(email string) *domain.Customer
	CreateCustomer(customer domain.Customer) domain.Customer
}
