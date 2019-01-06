package taxService

const(
  TAX_FOODANDBEVERAGE = 0.1 //equals with 10%
  TAX_TOBACCO = 0.02 // equals with (2/100)
  TAX_ENTERTAINMENT = 0.01 // equals with 1/100 and this tax will be free in case the price given less than 100

  TAX_FREE = 0
)

type TaxHelper interface{
    Calculate() float64
    Type() string
    Refundable() string
}

type Detail struct{
    Type string
    Refundable string
    Tax float64
    Amount float64
}

type Tobacco struct{
    detail Detail
}

type FoodAndBeverage struct{
    detail Detail
}

type Entertainment struct{
    detail Detail
}

type Default struct{
    detail Detail
}

func (tobacco Tobacco) Calculate() (tax float64){
    return TAX_TOBACCO
}

func (tobacco Tobacco) Type() (typeOfTax string){
    return "TOBACCO"
}

func (tobacco Tobacco) Refundable() (refundable string){
    return "NO"
}

func (foodAndBeverage FoodAndBeverage) Calculate() (tax float64){
    return TAX_FOODANDBEVERAGE
}

func (tobacco FoodAndBeverage) Type() (typeOfTax string){
    return "FOOD AND BEVERAGE"
}

func (tobacco FoodAndBeverage) Refundable() (refundable string){
    return "YES"
}

func (entertainment Entertainment) Calculate() (tax float64){
    return TAX_ENTERTAINMENT
}

func (entertainment Entertainment) Type() (typeOfTax string){
    return "ENTERTAINMENT"
}

func (entertainment Entertainment) Refundable() (refundable string){
    return "NO"
}

func CalculateHelper(t TaxHelper) (tax float64){
    return t.Calculate()
}

func TypeHelper(t TaxHelper) (typeOfTax string){
    return t.Type()
}

func RefundableHelper(t TaxHelper) (refundable string) {
    return t.Refundable()
}
