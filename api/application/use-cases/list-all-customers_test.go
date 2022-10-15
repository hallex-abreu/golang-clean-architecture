package usecases_test

import (
	"fmt"
	"strconv"
	"testing"

	usecases "github.com/hallex-abreu/golang-clean-architecture/api/application/use-cases"
	inmemory "github.com/hallex-abreu/golang-clean-architecture/api/infra/database/in-memory"
)

func Test_Get_All_Customers(t *testing.T) {
	var name = "Hallex Abreu"
	var email = "contato.hallex@gmail.com"
	var phone = "98383838383"

	inMemoryCustomer := inmemory.InMemoryCustomer{}

	for i := 0; i < 10; i++ {
		customerRequest := usecases.CreateCustomerRequest{
			Name:  name + strconv.Itoa(i),
			Email: email + strconv.Itoa(i),
			Phone: phone + strconv.Itoa(i),
		}

		usecases.CreateCustomer(customerRequest, inMemoryCustomer)
	}

	pagination := usecases.ListAllCustomersRequest{
		Page:   1,
		Limit:  10,
		Filter: "",
	}

	customers, _ := usecases.ListAllCustomers(pagination, inMemoryCustomer)
	inMemoryCustomer.Clear()

	if len(customers.Content) == 10 {
		fmt.Printf("\"list_all_customers\" SUCCEDED, expected -> %d, got -> %d", 10, len(customers.Content))
		t.Logf("\"list_all_customers\" SUCCEDED, expected -> %d, got -> %d", 10, len(customers.Content))
	} else {
		fmt.Printf("\"list_all_customers\" FAILED, expected -> %d, got -> %d", 10, len(customers.Content))
		t.Errorf("\"list_all_customers\" FAILED, expected -> %d, got -> %d", 10, len(customers.Content))
	}
}

func Test_Get_All_Customers_No_Return(t *testing.T) {
	inMemoryCustomer := inmemory.InMemoryCustomer{}

	pagination := usecases.ListAllCustomersRequest{
		Page:   1,
		Limit:  10,
		Filter: "",
	}

	_, err := usecases.ListAllCustomers(pagination, inMemoryCustomer)
	inMemoryCustomer.Clear()

	if err != nil {
		fmt.Printf("\"list_all_customers_no_return\" SUCCEDED, MENSAGE -> %s", err)
		t.Logf("\"list_all_customers_no_return\" SUCCEDED, MENSAGE -> %s", err)
	} else {
		fmt.Printf("\"list_all_customers_no_return\" FAILED")
		t.Errorf("\"list_all_customers_no_return\" FAILED")
	}
}
