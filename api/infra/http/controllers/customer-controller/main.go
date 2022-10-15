package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	usecases "github.com/hallex-abreu/golang-clean-architecture/api/application/use-cases"
	"github.com/hallex-abreu/golang-clean-architecture/api/domain"
	"github.com/hallex-abreu/golang-clean-architecture/api/infra/database/gorm/repositories"
)

func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	filter := c.DefaultQuery("filter", "")

	customerRepository := repositories.CustomerRepository{}

	pagination := usecases.ListAllCustomersRequest{
		Page:   page,
		Limit:  limit,
		Filter: filter,
	}

	customers, err := usecases.ListAllCustomers(pagination, customerRepository)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, customers)
	}
}

func Store(c *gin.Context) {
	var body domain.Customer

	customerRepository := repositories.CustomerRepository{}

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := usecases.CreateCustomerRequest{
		Name:  body.Name,
		Email: body.Email,
		Phone: body.Phone,
	}

	_, err = usecases.CreateCustomer(customer, customerRepository)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
