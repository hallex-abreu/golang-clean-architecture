package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	usecases "github.com/hallex-abreu/golang-clean-architecture/api/application/use-cases"
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
