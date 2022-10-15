package repositories

import (
	"github.com/hallex-abreu/golang-clean-architecture/api/domain"
	"github.com/hallex-abreu/golang-clean-architecture/api/infra/database/gorm"
)

type CustomerRepository struct{}

func (c CustomerRepository) FindCustomerByEmail(email string) *domain.Customer {
	var customer domain.Customer

	if err := gorm.DB.Where("email = ?", email).First(&customer).Error; err != nil {
		return nil
	} else {
		return &customer
	}
}

func (c CustomerRepository) FindAllCustomers(page int, size int, filter string) *[]domain.Customer {
	var customers []domain.Customer

	err := gorm.DB.Where("name LIKE ?", "%"+filter+"%").Or("email LIKE ?", "%"+filter+"%").Offset(size * (page - 1)).Limit(size).Find(&customers)

	if err != nil {
		return &customers
	} else {
		return nil
	}
}

func (c CustomerRepository) GetCountCustomers(filter string) int {
	var count int64
	var customerCount []domain.Customer
	gorm.DB.Where("name LIKE ?", "%"+filter+"%").Or("email LIKE ?", "%"+filter+"%").Find(&customerCount).Count(&count)
	return int(count)
}

func (c CustomerRepository) CreateCustomer(customer domain.Customer) domain.Customer {
	gorm.DB.Create(&customer)
	return customer
}
