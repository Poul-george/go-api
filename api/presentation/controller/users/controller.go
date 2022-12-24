package users

import (
	"github.com/labstack/echo/v4"
)

type Prameter struct {
	Name                 string `json:"name"`
	Comments             string `json:"comments"`
	MailAddress          string `json:"email"`
	Password             string `json:"password"`
	ConfirmationPassword string `json:"confirmationPassword"`
}

func Controller(c echo.Context) error {
	var p Prameter
	if err := c.Bind(&p); err != nil {
		return err
	}
	// fmt.Printf("--------------------{[%v]}----------------------------ooooo\n", c)
	// fmt.Printf("--------------------{[%v]}----------------------------ooooo\n", p)
	// fmt.Printf("--------------------{[%v]}----------------------------ooooo\n", p.Name)
	// fmt.Printf("--------------------{[%v]}----------------------------00000\n", p.Comments)
	// fmt.Printf("--------------------{[%v]}----------------------------ooooo\n", p.MailAddress)
	// fmt.Printf("--------------------{[%v]}----------------------------00000\n", p.Password)
	// fmt.Printf("--------------------{[%v]}----------------------------00000\n", p.ConfirmationPassword)
	return nil
}
