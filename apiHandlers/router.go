package apiHandlers

import (
	"chaithApp/api"
	"chaithApp/dbConfig"

	"context"


	"github.com/aws/aws-lambda-go/events"
)



func Router(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {
	dbConfig.ConnectToMongoDB()
	method := request.QueryStringParameters["method"]
	switch method {
case "CreateCustomer":
return api.CreateCustomerApi(ctx, request)

case "UpdateCustomer":
return api.UpdateCustomerApi(ctx, request)

case "DeleteCustomer":
return api.DeleteCustomerApi(ctx, request)

case "FindCustomer":
return api.FindCustomerApi(ctx, request)

case "FindallCustomer":
return api.FindallCustomerApi(ctx, request)

default:
return events.APIGatewayProxyResponse{StatusCode: 405, Body: "Method Not Allowed"},nil
}

}