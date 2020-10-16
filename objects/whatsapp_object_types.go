package objects

import (
	"fmt"
	"os"
	"strconv"
)

const (
	TwoFactorPin = 739253
)

var (
	Tiers = []string{ "TIER 1", "TIER 2", "TIER 3" }
	Types = []string{ "DEVELOPMENT", "PRODUCTION" }
	STATUS = []string{ "APPROVAL_PENDING", "ACCEPTED", "IN_PROGRESS", "VERIFICATION_PENDING", "ACTIVE", "REJECTED", "DELETED" }
	WaVersion = "v2.29.3"
)

type GenericResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type RequestError struct {
	Error string `json:"error"`
	Code int `json:"code"`
}

type WATemplate struct {
	Namespace string `json:namespace"`
	Content string `json:"content"`
	Name string `json:"name"`
	Params int8 `json:"params"`
	Code string `json:"code"`
}

type WAAccount struct {
	ClientName string `json:"clientName"`
	Namespace string `json:"namespace"`
	Number int64 `json:"number"`
	CountryCode int8 `json:"countryCode"`
	PhoneNumber int64 `json:"phoneNumber"`
	AccountType string `json:"accountType"`
	Ip string `json:"ip"`
	TwoFactorPin int8 `json:"twoFactorPin"`
	Certificate string `json:"certificate"`
	Version string `json:"version"`
	Templates []WATemplate `json:"templates"`
	Backup string `json:"backup"`
	SendMessageStatusEvent bool `json:"sendMessageStatusEvent"`
	Tier string `json:"tier"`
}

// Includes takes an array of string as input and returns a bool if the given value exits in the array
func Includes(arr []string, value string) bool {
	for _, item := range arr{
		if item == value{
			return true
		}
	}
	return false
}

// Validate provides a method to validates request body on the incoming request body
// it will also prefill the optional fields with the appropriate values
func (acc *WAAccount)Validate() *[]RequestError {
	errs := make([]RequestError, 0)
	switch {
	case acc.ClientName == "":
		errs = append(errs, RequestError{
			Error: "ClientName is a required field",
			Code:  1400,
		})
	case acc.Namespace == "":
		errs = append(errs, RequestError{
			Error: "Namespace is a required field",
			Code:  1400,
		})
	case acc.CountryCode == 0:
		errs = append(errs, RequestError{
			Error: "CountryCode is a required field",
			Code:  1400,
		})
	case acc.PhoneNumber == 0:
		errs = append(errs, RequestError{
			Error: "PhoneNumber is a required field",
			Code:  1400,
		})
	case acc.AccountType == "" || !Includes(Types, acc.AccountType):
		errs = append(errs, RequestError{
			Error: "AccountType is a required field",
			Code:  1400,
		})
	case acc.Tier == "" || !Includes(Tiers, acc.Tier):
		errs = append(errs, RequestError{
			Error: "Namespace is a required field",
			Code:  1400,
		})
	case acc.Certificate == "":
		errs = append(errs, RequestError{
			Error: "Certificate is a required field",
			Code:  1400,
		})
	case acc.Tier == "" || !Includes(Tiers, acc.Tier):
		errs = append(errs, RequestError{
			Error: "Tier is a required field",
			Code:  1400,
		})
	case acc.TwoFactorPin == 0:
		acc.TwoFactorPin = TwoFactorPin
	case acc.Number == 0:
		num, _ := strconv.Atoi(fmt.Sprintf("%d%d", acc.CountryCode, acc.PhoneNumber))
		acc.Number = int64(num)
	case acc.Version == "":
		acc.Version = WaVersion
		if v := os.Getenv("WA_VERSION"); v != ""{
			acc.Version = v
		}
	}
	return &errs
}

// ValidateTemplate provides a method to validate templates, if not matched returns an array of RequestError
func (acc *WAAccount)ValidateTemplate() *[]RequestError {
	errs := make([]RequestError, 0)
	for _, tem := range acc.Templates{
		switch {
		case tem.Name == "":
			errs = append(errs, RequestError{
				Error: "Template.Name is a required field",
				Code:  1400,
			})
		}
	}
	return &errs
}