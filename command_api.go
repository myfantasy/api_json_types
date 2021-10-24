package api_json_types

import (
	"context"
	"encoding/json"

	"github.com/myfantasy/mft"
)

type CommandRequest struct {
	ObjectName  string          `json:"object,omitempty"`
	CommandName string          `json:"name,omitempty"`
	Params      json.RawMessage `json:"params,omitempty"`
	User        string          `json:"user,omitempty"`
}

type CommandResponce struct {
	Error  *mft.Error      `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

func (cr *CommandRequest) UserName() string {
	return cr.User
}

func (cr *CommandRequest) Unmarshal(v interface{}) *mft.Error {
	er0 := json.Unmarshal(cr.Params, v)
	if er0 != nil {
		return mft.GenerateErrorE(20000000, er0).
			AppendLabel(ErrorType, ErrorTypeInvalidParams)
	}

	return nil
}

func CreateRequest(objectName string, commandName string, v interface{}) (cr *CommandRequest, err *mft.Error) {
	b, er0 := json.Marshal(v)
	if er0 != nil {
		return nil, mft.GenerateErrorE(20000010, er0)
	}

	return &CommandRequest{
		ObjectName:  objectName,
		CommandName: commandName,
		Params:      b,
	}, nil
}

func CreateUserRequest(objectName string, commandName string, v interface{}, userName string) (cr *CommandRequest, err *mft.Error) {
	b, er0 := json.Marshal(v)
	if er0 != nil {
		return nil, mft.GenerateErrorE(20000011, er0)
	}

	return &CommandRequest{
		ObjectName:  objectName,
		CommandName: commandName,
		Params:      b,
		User:        userName,
	}, nil
}

func (cr *CommandResponce) Unmarshal(v interface{}) *mft.Error {
	er0 := json.Unmarshal(cr.Result, v)
	if er0 != nil {
		return mft.GenerateErrorE(20000020, er0).
			AppendLabel(ErrorType, ErrorTypeInternalError)
	}

	return nil
}

func CreateResponce(v interface{}, err *mft.Error) (cr *CommandResponce) {
	if err != nil {
		return &CommandResponce{
			Error: err,
		}
	}
	if v == nil {
		return &CommandResponce{}
	}
	b, er0 := json.Marshal(v)
	if er0 != nil {
		return &CommandResponce{
			Error: mft.GenerateErrorE(20000030, er0).
				AppendLabel(ErrorType, ErrorTypeInternalError),
		}
	}

	return &CommandResponce{
		Result: b,
	}
}

type CommandDescription struct {
	ObjectName  string `json:"object_name,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Api interface {
	AllowedCommands() []CommandDescription
	DoRequest(ctx context.Context, req *CommandRequest) *CommandResponce
}
