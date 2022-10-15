package repositories

import "github.com/hallex-abreu/golang-clean-architecture/api/domain"

type CustomerRepository interface {
	FindCustomerByEmail(email string) *domain.Customer
	FindAllCustomers(page int, size int, filter string) *[]domain.Customer
	GetCountCustomers(filter string) int
	CreateCustomer(customer domain.Customer) domain.Customer
}
