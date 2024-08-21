package customEcho

import "github.com/labstack/echo/v4"

type Context struct {
	echo.Context
}

func (c *Context) BindAndValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	return GetValidator().Struct(i)
}

func ContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &Context{c}
		return next(cc)
	}
}
