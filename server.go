package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Hello, World!
	// Browse to http://localhost:1323 and you should see Hello, World! on the page.
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Path Parameters
	// Browse to http://localhost:1323/users/Joe and you should see ‘Joe’ on the page.
	e.GET("/users/:id", func(c echo.Context) error {
		// User ID from path `users/:id`
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})

	// Query Parameters
	// Browse to http://localhost:1323/show?team=x-men&member=wolverine and you should see ‘team:x-men, member:wolverine’ on the page.
	e.GET("/show", func(c echo.Context) error {
		// Get team and member from the query string
		team := c.QueryParam("team")
		member := c.QueryParam("member")
		return c.String(http.StatusOK, "team:"+team+", member:"+member)
	})

	// Form application/x-www-form-urlencoded
	// curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/save
	e.POST("/save", func(c echo.Context) error {
		// Get name and email
		name := c.FormValue("name")
		email := c.FormValue("email")
		return c.String(http.StatusOK, "name:"+name+", email:"+email)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
