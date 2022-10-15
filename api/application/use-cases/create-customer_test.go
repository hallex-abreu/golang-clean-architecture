package usecases_test

import (
	"fmt"
	"testing"

	usecases "github.com/hallex-abreu/golang-clean-architecture/api/application/use-cases"
	"github.com/hallex-abreu/golang-clean-architecture/api/domain"
	inmemory "github.com/hallex-abreu/golang-clean-architecture/api/infra/database/in-memory"
)

func Test_Should_be_create_an_customer(t *testing.T) {
	var name = "Hallex Abreu"
	var email = "contato.hallex@gmail.com"
	var phone = "98383838383"

	inMemoryCustomer := inmemory.InMemoryCustomer{}

	customerRequest := usecases.CreateCustomerRequest{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	customer, _ := usecases.CreateCustomer(customerRequest, inMemoryCustomer)
	inMemoryCustomer.Clear()

	if customer.Name == customerRequest.Name {
		fmt.Printf("\"create_customer\" SUCCEDED, expected -> %s, got -> %s", name, customer.Name)
		t.Logf("\"create_customer\" SUCCEDED, expected -> %s, got -> %s", name, customer.Name)
	} else {
		fmt.Printf("\"create_customer\" FAILED, expected -> %s, got -> %s", name, customer.Name)
		t.Errorf("\"create_customer\" FAILED, expected -> %s, got -> %s", name, customer.Name)
	}
}

func Test_Should_be_validation_customer_with_email_already_in_use(t *testing.T) {
	var name = "Hallex Abreu"
	var email = "contato.hallex@gmail.com"
	var phone = "98383838383"

	inMemoryCustomer := inmemory.InMemoryCustomer{}

	customerRequest := usecases.CreateCustomerRequest{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	inMemoryCustomer.CreateCustomer(domain.Customer{
		Name:  name,
		Email: email,
		Phone: phone,
	})

	_, err := usecases.CreateCustomer(customerRequest, inMemoryCustomer)
	inMemoryCustomer.Clear()

	if err != nil {
		fmt.Printf("\"create_customer_with_email_already_in_use\" SUCCEDED, MENSAGE -> %s", err)
		t.Logf("\"create_customer_with_email_already_in_use\" SUCCEDED, MENSAGE -> %s", err)
	} else {
		fmt.Printf("\"create_customer_with_email_already_in_use\" FAILED")
		t.Errorf("\"create_customer_with_email_already_in_use\" FAILED")
	}
}
