package model

import (
    "fmt"
    "go-tax/conf"
)

type Tax struct{
    Name string
    Code int
    Price int
    CreatedAt string
    UpdatedAt string
}

const(
    TYPE_FOODANDBEVERAGE = 1
    TYPE_TOBACCO = 2
    TYPE_ENTERTAINMENT = 3
)

func TaxModelStore(tax *Tax) (status bool){

    //connect to database
    db, err := conf.Connect()
    if err != nil {
        fmt.Println("Error to connect : ", err.Error())
        return false
    }

    defer db.Close()

    //insert
    _, err = db.Exec("insert into taxes(`name`, `code`, `price`, `created_at`, `updated_at`) values (?, ?, ?, ?, ?)", &tax.Name, &tax.Code, &tax.Price, &tax.CreatedAt, &tax.UpdatedAt)

    if err != nil {
        fmt.Println("Error to insert database : ", err.Error())
        return false
    }

    return true
}

func TaxModelGet() (taxes []Tax){

  db, err := conf.Connect()
  if err != nil {
      fmt.Println("Error to connect : ", err.Error())
      return
  }

  defer db.Close()
  rows, err := db.Query("select name, code, price, created_at, updated_at from taxes")

  for rows.Next() {
      var each = Tax{}

      var err = rows.Scan(&each.Name, &each.Code, &each.Price, &each.CreatedAt, &each.UpdatedAt)

      if err != nil {
          fmt.Println(err.Error())
          return
      }

      taxes = append(taxes, each)

  }

  return taxes
}
