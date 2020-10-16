package rest_handlers

import (
	"encoding/json"
	"fmt"
	"github.com/parvez0/whatsapp-provisioner/objects"
	"io"
	"io/ioutil"
	"net/http"
)

// CreateAccount stores the new infra request to sqlite db
func CreateAccount(writer http.ResponseWriter, request *http.Request) {
	defer RecoverFromPanic(&writer)
	reqBody, err := ioutil.ReadAll(request.Body)
	genResp := objects.GenericResponse{
		Success: false,
		Message: "Bad Request",
	}
	if err != nil || err == io.EOF {
		data := objects.WAAccount{}
		genResp.Data = data.Validate()
		BadRequest(&writer, &genResp)
		return
	}
	body := objects.WAAccount{}
	err =json.Unmarshal(reqBody, &body)
	if err != nil{
		genResp.Data = []objects.RequestError{
			{
				Error: fmt.Sprintf("failed to parse body: %v", err),
				Code:  1400,
			},
		}
		BadRequest(&writer, &genResp)
		return
	}
	errs := body.Validate()
	if len(*errs) > 0{
		genResp.Data = errs
		BadRequest(&writer, &genResp)
		return
	}
	genResp.Success = true
	genResp.Message = "Request accepted"
	resp, _ := json.Marshal(genResp)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(201)
	writer.Write(resp)
}