# golang-sample-api
Sample API Using GoLang gin-gonic Framework


## Prerequisites
1. Go
2. PostgreSQL
3. Text Editor/IDE. Recommendation: GoLand
4. GOPATH setup

## Package
1. github.com/gin-gonic/gin : GoLang HTTP web framework
2. github.com/go-sql-driver/mysql : GoLang MySQL-Driver
3. github.com/jinzhu/gorm : GoLang ORM library
4. github.com/spf13/viper : Golang configuration library to read your application environment variables

run `go get packageName` to install package

## Configraton File
Create new file `config.yml` in root directory. File Content:
```
app:
  host: yourApplicationHost e.g. http://localhost:8080
  port: yourApplicationPort e.g. 8080
  show-sql: true
  production: false
datasource: "user:password@/db_name?charset=utf8&parseTime=True&loc=Local"
```

## Database Table
Create new Table on your database. You need to run it manually because we don't use migration process in this application.

```
CREATE TABLE IF NOT EXISTS user_users (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(50) NOT NULL,
  skin_type varchar(50),
  age int(11),
  updated_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

 CREATE TABLE IF NOT EXISTS user_skin_cares (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  brand varchar(50),
  name varchar(50) NOT NULL,
  product_type varchar(50),
  skin_concern varchar(100),
  updated_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  user_id int not null,
   FOREIGN KEY fk_user_skin_cares_user_users_user_id(user_id)
   REFERENCES user_users(id)
   ON UPDATE CASCADE
   ON DELETE RESTRICT
);
```

## How To Run
1. Run `go build -o sample-golang-api` to build Application
2. Then use `./sample-golang-api` to run Application

if your application run successfully, here is the result :
```
[GIN-debug] GET    /v1/user/user             --> github.com/yaumianwar/sample-golang-api/user/handler.GetAllUsers.func1 (3 handlers)
[GIN-debug] POST   /v1/user/user             --> github.com/yaumianwar/sample-golang-api/user/handler.CreateUser.func1 (3 handlers)
[GIN-debug] GET    /v1/user/user/:id         --> github.com/yaumianwar/sample-golang-api/user/handler.GetSingleUser.func1 (3 handlers)
[GIN-debug] POST   /v1/user/user/:id         --> github.com/yaumianwar/sample-golang-api/user/handler.UpdateUser.func1 (3 handlers)
[GIN-debug] DELETE /v1/user/user/:id         --> github.com/yaumianwar/sample-golang-api/user/handler.DeleteUser.func1 (3 handlers)
[GIN-debug] GET    /v1/user/user/:id/skincare --> github.com/yaumianwar/sample-golang-api/user/handler.GetAllSkincareByUser.func1 (3 handlers)
[GIN-debug] POST   /v1/user/user/:id/skincare --> github.com/yaumianwar/sample-golang-api/user/handler.CreateSkincare.func1 (3 handlers)
[GIN-debug] GET    /v1/user/user/:id/skincare/:skincareId --> github.com/yaumianwar/sample-golang-api/user/handler.GetSingleSkincare.func1 (3 handlers)
[GIN-debug] POST   /v1/user/user/:id/skincare/:skincareId --> github.com/yaumianwar/sample-golang-api/user/handler.UpdateSkincare.func1 (3 handlers)
[GIN-debug] DELETE /v1/user/user/:id/skincare/:skincareId --> github.com/yaumianwar/sample-golang-api/user/handler.DeleteSkincare.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```
## Summary API
This application is about user skincare, this application allow us to :
1. Get All User Data : API `GET /v1/user/user`
2. Get Single User Data : API `GET /v1/user/user/:id`
3. Create User Data : API `POST /v1/user/user`
4. Update user Data : API `POST /v1/user/user/:id`
5. Delete User Data : API `DELETE /v1/user/user/:id`
6. Get All Skincare Data By User :  API `GET /v1/user/user/:id/skincare`
7. Get Single Skincare Data : API `GET /v1/user/user/:id/skincare/:skincareId`
8. Create User Skincare Data : API `POST /v1/user/user/:id/skincare`
9. Update User Skincare Data : API `POST /v1/user/user/:id/skincare/:skincareId`
10. Delete User Skincare Data : API `DELETE /v1/user/user/:id/skincare/:skincareId`

## API Request
After your application already running, you can create API Request using `curl` or Postman. 
1. I provide Postman Collection at root directory called `sample-golang-api.postman_collection.json`, so you can import this file to your Postman Application and will create new Postman Collection for this application API which contains all of API of this application
2. `curl` example to request API Get All User Data
    
    run :
    
    `curl -X GET http://localhost:8080/v1/user/user`
    
    result :
    ```
    {"data":[{"id":1,"name":"Nur Yaumi","skinType":"oily acne prone","age":20,"updatedAt":"2019-07-16T19:41:27+08:00","createdAt":"2019-07-16T19:41:27+08:00"}],"status":{"message":"data fetched"}}
    ```
