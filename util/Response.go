package util

import(
    f "fmt"
    "encoding/json"
    "net/http"
)

type Response struct{
    Code int
    MessageCode string
    Message string
    Data *interface{}
}

func ComposeResponse(w http.ResponseWriter, i interface{}, messageCode string, message string)  {

    var response Response

    switch messageCode {
    case "OK":
        response.Code = 200
        response.MessageCode = messageCode
        response.Message = message
        response.Data = &i
    case "204":
        f.Println("NO CONTENT!")
    case "400":
        f.Println("BAD REQUEST!")
    case "401":
        f.Println("UNAUTHORIZED!")
    case "403":
        f.Println("FORBIDDEN!")
    case "405":
        f.Println("METHOD NOT ALLOWED!")
    case "UNPROCESSABLE ENTITY":
      response.Code = 422
      response.MessageCode = messageCode
      response.Message = message
      response.Data = nil
    default:
      response.Code = 400
      response.MessageCode = "BAD REQUEST"
      response.Message = "Something was Error!"
      response.Data = nil
    }

    res, err := json.Marshal(response)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
    w.Write([]byte("\n"))

}
