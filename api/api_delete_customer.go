package api

import (
  "chaithApp/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  )


// @Summary      DELETE Customer input: Customer
// @Description  DELETE Customer API
// @Tags         DELETE Customer - Customer
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Customer "Status OK"
// @Success      202  {array}   dto.Model_Customer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteCustomer [DELETE]


    func DeleteCustomerApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


objectid := event.QueryStringParameters["objectid"]
    err := dao.DB_DeleteCustomer(objectid)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       string(""),
    },nil
    

}
