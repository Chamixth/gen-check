package api

import (
  "chaithApp/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  
  "encoding/json"
)


// @Summary      GET Customer input: Customer
// @Description  GET Customer API
// @Tags         GET Customer - Customer
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Customer "Status OK"
// @Success      202  {array}   dto.Model_Customer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindCustomer [GET]


    func FindCustomerApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


objectid := event.QueryStringParameters["objectid"]
    inputObj, err := dao.DB_FindCustomer(objectid)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


responseBodyJSON, _ := json.Marshal(inputObj)

    return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBodyJSON),
	},nil

}
