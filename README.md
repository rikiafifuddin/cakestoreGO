# CAKE STORE

- RESTFul API
- Full-featured: GET List of Cake, GET Detail of Cake, ADD New Cake, Update Cake (PATCH), and DELETE Cake
- Include Unit Test
- Easy For modify
- Applied request validation

# Tools
> [GO](https://go.dev/) "Programing Language"
> 
> [MySQL](https://www.mysql.com/) "Database"
> 
> [Docker](https://www.docker.com/) "Container"

# Instalation

#### Download or Clone Repository

`$ git clone https://github.com/rikiafifuddin/cakestoreGO.git`

#### Database
open MySQL and create Database
 ###### Option 1 Import Database via import file
 - Open MySQL Workbench
 - Menu "Sever"-> "Data Import"
 - Import from Self-contained file
 - select `cakestore_cake.sql` provided
 - chose default target schema
 - start import

###### Option 2 Import Database via SQL script
- run SQL script provide `database.sql`

# RUNING
 ###### Option 1 runing program via terminal
 - open your directory
 - `$ cd .\api\`
 - `$ go run main.go`
 - use API path on documentation or use Postman Collection provided `PrivyID.postman_collection.json` (import to your postman app)
 
###### Option 2 runing program via Docker
- make sure you have docker instaled on your device
- open terminal `$ docker build -t [container name] . ` 
example `$ docker build -t privyid . `
- wait until docker finish building container
- run docker `$ docker run -p 8080:8080 -t [container name]`
example `$ docker run -p 8080:8080 -t privyid`
- use API path on documentation or use Postman Collection provided `PrivyID.postman_collection.json` (import to your postman app)

#### PATH　

```go
	router.HandleFunc("/cakes", service.ListOfCake).Methods("GET")
	router.HandleFunc("/cakes/{id}", service.DetailOfCake).Methods("GET")
	router.HandleFunc("/cakes", service.AddNewCake).Methods("POST")
	router.HandleFunc("/cakes/{id}", service.UpdateCake).Methods("PATCH")
	router.HandleFunc("/cakes/{id}", service.DeleteCake).Methods("DELETE")
```

#### JSON BODY REQUEST FOR ADD NEW CAKE

```json
{
    "title" : "Maple",
    "description" : "Cake From heaven",
    "rating" : 7,
    "Image" : "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
}
```
> all variable must be fill except "rating"

#### JSON BODY REQUEST FOR PATCH or UPDATE CAKE
```json
{
	"description" : "Cake From heaven",
    	"rating" : 7,
}
```
> Edit/patch whatever you want to edit, one or more variable at the same time

#### THANKS
`IF YOU HAVE QUESTION YOU CAN CONTACT ME ON EMAIL: rikinur34@gmail.com`
> "Technical Test LINK" [GO Backend Engineer](https://peach-advantage-0ff.notion.site/Technical-Test-Backend-Engineer-632ff30b3e854de7a7fb859fccf98f19) 

![](https://pandao.github.io/editor.md/examples/images/4.jpg)
