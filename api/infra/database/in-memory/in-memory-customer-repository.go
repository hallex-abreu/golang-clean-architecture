package inmemory

import "github.com/hallex-abreu/golang-clean-architecture/api/domain"

type InMemoryCustomer struct{}

var items = []domain.Customer{}

func (i InMemoryCustomer) FindCustomerByEmail(email string) *domain.Customer {
	for _, c := range items {
		if email == c.Email {
			return &c
		}
	}

	return nil
}

func (i InMemoryCustomer) FindAllCustomers(page int, size int, filter string) *[]domain.Customer {
	return &items
}

func (i InMemoryCustomer) GetCountCustomers(filter string) int {
	return len(items)
}

func (i InMemoryCustomer) CreateCustomer(customer domain.Customer) domain.Customer {
	items = append(items, customer)
	return customer
}

// Used for clear items in end of unit test
func (i InMemoryCustomer) Clear() {
	items = items[:0]
}
