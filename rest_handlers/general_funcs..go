package rest_handlers

import (
	"encoding/json"
	"fmt"
	"github.com/parvez0/whatsapp-provisioner/custom_logger"
	"github.com/parvez0/whatsapp-provisioner/objects"
	"net/http"
)

var clog = custom_logger.NewLogger()

// RecoverFromPanic will provide a generic function for handling unwanted panic during rest call
func RecoverFromPanic(writer *http.ResponseWriter)  {
	if err := recover(); err != nil {
		genResp := objects.GenericResponse{
			Success: false,
			Message: fmt.Sprintf("%v", err),
			Data:    nil,
		}
		(*writer).Header().Set("Content-Type", "application/json")
		(*writer).WriteHeader(500)
		resp, err := json.Marshal(genResp)
		if err != nil {
			(*writer).Write([]byte("internal error"))
			return
		}
		(*writer).Write(resp)
	}
}

// BadRequest responds with bad request error
func BadRequest(writer *http.ResponseWriter, data *objects.GenericResponse)  {
	(*writer).Header().Set("Content-Type", "application/json")
	(*writer).WriteHeader(400)
	resp, err := json.Marshal(data)
	if err != nil{
		clog.Errorf("failed to marshal json object %+v - error - %+v", data, err)
		(*writer).Write([]byte("one or more required parameters not provided"))
		return
	}
	(*writer).Write(resp)
	return
}