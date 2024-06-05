package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Path    string      `json:"path,omitempty"`
	Message string      `json:"message,omitempty"`
	Reason  string      `json:"reason,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func WrapWithCode(err error, code int, message string) *WebResponse {
	return &WebResponse{
		Code:    code,
		Status:  http.StatusText(code),
		Message: message,
		Reason:  err.Error(),
	}
}

func (ve *WebResponse) Error() string {
	return fmt.Sprintf("value error: %s", ve.Reason)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func FatalIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}

// func CommitOrRollback(tx *sql.Tx) {
// 	err := recover()
// 	if err != nil {
// 		errorRollback := tx.Rollback()
// 		PanicIfError(errorRollback)
// 		panic(err)
// 	} else {
// 		errorCommit := tx.Commit()
// 		PanicIfError(errorCommit)
// 	}
// }
