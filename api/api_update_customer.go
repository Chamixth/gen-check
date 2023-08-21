package api

import (
  "chaithApp/dao"
  "context"
  "github.com/aws/aws-lambda-go/events"
  "net/http"

  
    "encoding/json"
    "chaithApp/dto"

  )


// @Summary      PUT Customer input: Customer
// @Description  PUT Customer API
// @Tags         PUT Customer - Customer
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Customer false "string collection" 
// @Success      200  {array}   dto.Model_Customer "Status OK"
// @Success      202  {array}   dto.Model_Customer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateCustomer [PUT]


    func UpdateCustomerApi(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {


inputObj := dto.Customer{}
    if err := json.Unmarshal([]byte(event.Body), &inputObj); err != nil {
    		return  events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "Invalid Request Body"},nil
    }

    //validate := validator.New()
    //if validationErr := validate.Struct(&inputObj); validationErr != nil {
      //  return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    //}
err := dao.DB_UpdateCustomer(&inputObj)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()},nil
    }


responseBodyJSON, _ := json.Marshal(inputObj)
    return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBodyJSON),
	},nil
    
    

}
