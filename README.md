# Golang boiler-plate with clean architecture
### Author: Hieuhoccode
Mục đích nhăm giúp các bạn mới bắt đầu với golang có thể hiểu được cách:
- cấu trúc một project golang
- cách viết một API đơn giản với golang
- cách kết nối với database
- Link bài viết chi tiết: [ở đây](https://t.ly/clbnD)
## need to install before run
- [golang](https://golang.org/doc/install)
- [docker](https://docs.docker.com/get-docker/)
- postgresql
- dbeaver or pgadmin
- postman
- goland or vscode

## source code structure
```
├── README.md
├── geometry # bài tập về hình học
├── backend-basic # viết API và connect database trong một file
├── backend-refactor # tái cấu trúc code backend-basic
    ├── api # api mà gọi đến service khác
    ├── route # api mà frontend sẽ gọi
    ├── handler # xử lý các request và trả về response
    ├── service # logic chính của ứng dụng
    ├── repository # tương tác với database
    ├── config # các file config
        ├── app # config app
        ├── db # config database
        ├── env # chứa các biến môi trường
    ├── middleware # các hàm trung gian, xử lý yêu cầu trước khi nó tới handler
    ├── model # lưu các struct
    ├── util # chứa các hàm tiện ích
        ├── constant.go # chứa các constant
        ├── error.go # chứa các error
        ├── response.go # chứa các response
        ├── pointer.go # chứa các hàm kiểm tra pointer
    ├── main.go # file chạy chính
    ├── go.mod # file quản lý các package
    ├── go.sum # file quản lý các package
    ├── .gitignore # file quản lý các file không cần push lên git
```

## package used
- github.com/gin-gonic/gin
- gorm.io/gorm
- gorm.io/driver/postgres
- github.com/caarlos0/env
- errors
- context
- fmt

## how to run
- clone project
- run command `go run main.go` to run project
- open postman and import file `postman_collection.json` to test API
