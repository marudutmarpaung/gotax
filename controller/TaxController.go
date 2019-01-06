package controller

import(
    "net/http"
    "go-tax/model"
    "strconv"
    "time"
    "go-tax/util"
    "go-tax/taxService"
    "fmt"
)

func TaxControllerStore(w http.ResponseWriter, r *http.Request)  {

    name := r.FormValue("name")

    code, err := strconv.Atoi(r.Form.Get("code"))
    if err != nil{
      fmt.Println("Failed to convert : ", err)
    }

    price, err := strconv.Atoi(r.Form.Get("price"))
    if err != nil{
      fmt.Println("Failed to convert : ", err)
    }

    tax := model.Tax{Name: name, Code:code , Price:price , CreatedAt:time.Now().Format("2006-01-02 15:04:05") , UpdatedAt: time.Now().Format("2006-01-02 15:04:05")}
    insert := model.TaxModelStore(&tax)

    if !insert{
        util.ComposeResponse(w, tax, "UNPROCESSABLE ENTITY", "Failed to Store Data!")
    }

    util.ComposeResponse(w, tax, "OK", "Success to Store Data!")

}

func TaxControllerGet(w http.ResponseWriter, r *http.Request)  {
    taxes := model.TaxModelGet()

    /*
      since data will be diffrent to be passed, need to define local struct
      the defined struct will be passed and consumed by Front End
    */
    type Result struct{
        Name string
        Code int
        Price int
        Detail taxService.Detail
    }

    var result Result
    var list []Result
    var tobacco taxService.Tobacco
    var foodAndBeverage taxService.FoodAndBeverage
    var entertainment taxService.Entertainment

    tax := taxService.TaxHelper(foodAndBeverage) // set as default

    for _, each := range taxes {
        result.Name = each.Name
        result.Code = each.Code
        result.Price = each.Price

        //clasify tax by its type in loop
        switch each.Code {
          case model.TYPE_FOODANDBEVERAGE:
              tax = taxService.TaxHelper(foodAndBeverage)
              result.Detail.Tax = taxService.CalculateHelper(tax) * float64(each.Price)
          case model.TYPE_TOBACCO:
              tax = taxService.TaxHelper(tobacco)
              result.Detail.Tax = 10 + taxService.CalculateHelper(tax) * float64(each.Price)
          case model.TYPE_ENTERTAINMENT:
              tax = taxService.TaxHelper(entertainment)
              if each.Price < 100 {
                  result.Detail.Tax = taxService.TAX_FREE
              }else{
                  // sum:= float64(each.Price - 100)
                  result.Detail.Tax = taxService.CalculateHelper(tax) * float64(each.Price - 100)
              }
        }

        result.Detail.Amount = (float64(each.Price) + result.Detail.Tax)

        result.Detail.Type = taxService.TypeHelper(tax)
        result.Detail.Refundable = taxService.RefundableHelper(tax)

        list = append(list, result)
    }

    util.ComposeResponse(w, list, "OK", "Success to get List Of Taxes!")
}
