package model

import (
    f "fmt"
    "go-tax/conf"
)

type Park struct{
    Id int64
    RegistrationId string
    Plat string
    TimeIn string
    TimeOut string
    Price int
    CreatedAt string
    UpdatedAt string
}

func ParkModelStore(registrationId, platNum, registerTime, currentTime string) (status bool){

    //connect to database
    db, err := conf.Connect()
    if err != nil {
        f.Println("Error to connect : ", err.Error())
        return
    }

    defer db.Close()

    //insert
    _, err = db.Exec("insert into parks(`registration_id`, `plat`, `time_in`, `created_at`, `updated_at`) values (?, ?, ?, ?, ?)", registrationId, platNum, registerTime, currentTime, currentTime)

    if err != nil {
        f.Println("Error to insert database : ", err.Error())
        return false
    }

    return true
}

func ParkModelGetByPlat(plat string) (park Park){
    //connect to database
    db, err := conf.Connect()
    if err != nil {
        f.Println("Error to connect : ", err.Error())
        return
    }

    defer db.Close()


    err = db.
        QueryRow("select * from parks where plat = ? and time_out is NULL", plat).
        Scan(&park.Id, &park.RegistrationId, &park.Plat, &park.TimeIn, &park.TimeOut, &park.Price, &park.CreatedAt, &park.UpdatedAt)
    if err != nil {
        f.Println(err.Error())
        return
    }

    return park
}

func ParkModelUpdatePrice(time_out string, updated_at string, price int, park Park) (status bool) {
    //connect to database
    db, err := conf.Connect()
    if err != nil {
        f.Println("Error to connect : ", err.Error())
        return
    }

    defer db.Close()

    _, err = db.Exec("update parks set time_out = ?, updated_at = ?, price = ? where id = ?", time_out, updated_at, price, park.Id)
    if err != nil {
        f.Println(err.Error())
        return false
    }

    return true
}

func ParkModelGetParkByDuration(start_time, end_time string) (parks []Park) {
    //connect to database
    db, err := conf.Connect()
    if err != nil {
        f.Println("Error to connect : ", err.Error())
        return
    }

    defer db.Close()
    rows, err := db.Query("select * from parks where `time_out` BETWEEN ? and ? ", start_time, end_time)

    for rows.Next() {
        var each = Park{}

        var err = rows.Scan(&each.Id, &each.RegistrationId, &each.Plat, &each.TimeIn, &each.TimeOut, &each.Price, &each.CreatedAt, &each.UpdatedAt)

        if err != nil {
            f.Println(err.Error())
            return
        }

        parks = append(parks, each)

    }

    return parks
}
