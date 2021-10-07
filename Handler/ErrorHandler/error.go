package ErrorHandler

import (
	"encoding/json"
	"net/http"
)


func Response(w http.ResponseWriter,status int,message string,code string){
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"Status": "Error occur while processing your request.",
		"Error ": message,
		"Code  ": code,
	})
}
func Response3(w http.ResponseWriter,status int,message string,code string){
	w.WriteHeader(status)
	type ErrorMessage struct {
		ErrorMessage string `json:"error"`
		Code         string `json:"code"`
	}

	json.NewEncoder(w).Encode(map[string]string{
		"Status": "Error occur while processing your request.",
		"Error ": message,
		"Code  ": code,
	})
}

func Response1(w http.ResponseWriter,status int,message string,product int,name string){
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"Status    ": "Your request processed Successfully.",
		"Message   ": message,
		"Product Id": (string)(product),
		"ProductName      ": name,
	})
}

func Response2(w http.ResponseWriter,status int,message string,product int,name string,total int){
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"Status    ": "Your request processed Successfully.",
		"Message   ": message,
		"Product Id": (string)(product),
		"ProductName      ": name,
		"Total bill": string(total),
	})
}
