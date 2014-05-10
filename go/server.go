package main

import "github.com/go-martini/martini"

func main() {
  m:= martini.Classic()
  //magic here
  m.Get("/hello", func() string {
    return "hello"
  })

  // m.Get("/response", func(res http.ResponseWriter, req *http.Request) {
  //   res.WriteHeader(200)
  // })
  // m.Post("/product" func() {
  //   //do some post here
  // })

  // m.Put("/product", func() {
  //   //update here
  // })

  // m.Delete("/product/:id", func() {
  //   //delete a product
  // })

  m.Get("/size/:hello", func(params martini.Params) string {
    return "Hello " + params["hello"]
  })
  //REST :)
  m.Group("/books", func(r martini.Router) {
    r.Get("/:id", GetBooks)
    r.Post("/new", NewBook)
    r.Put("/update/:id", UpdateBook)
    r.Delete("/delete/:id", DeleteBook)
})
  m.Run()

}
