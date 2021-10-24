package apijsontypes

import "github.com/myfantasy/mft"

const ErrorType mft.ErrorLabelName = "error_type"
const ErrorTypeInvalidParams string = "invalid_params"
const ErrorTypeInternalError string = "internal_error"

// Errors codes and description
var Errors map[int]string = map[int]string{
	20000000: "api_json_types.CommandRequest.Unmarshal: fail",
	20000010: "api_json_types.CreateRequest: Marshal fail",
	20000011: "api_json_types.CreateUserRequest: Marshal fail",
	20000020: "api_json_types.CommandResponce.Unmarshal: fail",
	20000030: "api_json_types.CreateResponce: fail",
}

func init() {
	mft.AddErrorsCodes(Errors)
}
