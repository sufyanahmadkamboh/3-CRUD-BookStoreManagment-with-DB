# 3-CRUD-BookStoreManagment-with-DB

go mod init github.com/sufyanahmadkamboh/3-CRUD-BookStoreManagment-with-DB
go get "github.com/gorilla/mux"
go get "gorm.io/gorm"
go get "gorm.io/driver/postgres"
go get "gorm.io/driver/mysql"

create table Books with 3 fields Name, Author, Publication