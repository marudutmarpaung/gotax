package conf

import(
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
)

const(
    DB_CONNECTION = "mysql"
    DB_HOST = "127.0.0.1"
    DB_PORT = "3306"
    DB_DATABASE = "parking"
    DB_USERNAME = "root"
    DB_PASSWORD = "Marudutif10017#dana"
)

/*
  connect into database
*/
func Connect() (*sql.DB, error) {
    config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)

    db, err := sql.Open(DB_CONNECTION, config)
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }

    return db, nil
}
