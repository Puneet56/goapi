package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func HandleUserGroup(g *echo.Group) {
	g.GET("", getUsers)
	g.GET("/:id", getUser)
	g.POST("", createUser)
	g.PUT("/:id", editUser)
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = make([]User, 0)

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	for _, user := range users {
		if user.ID == id {
			return c.JSON(http.StatusOK, user)
		}
	}

	return c.JSON(http.StatusNotFound, nil)
}

func createUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	user.ID = uuid.NewString()
	users = append(users, user)
	return c.JSON(http.StatusCreated, user)
}

func editUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	for _, u := range users {
		if u.ID == user.ID {
			u.Name = user.Name
			u.Email = user.Email

			return c.JSON(http.StatusOK, u)
		}
	}

	return c.JSON(http.StatusNotFound, "User not found"+user.ID)
}
