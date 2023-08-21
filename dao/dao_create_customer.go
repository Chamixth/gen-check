package dao

import (
    "context"
	"chaithApp/dbConfig"
	"chaithApp/dto"

)

func DB_CreateCustomer (application *dto.Customer) error {

	_, err := dbConfig.DATABASE.Collection("Customers").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}