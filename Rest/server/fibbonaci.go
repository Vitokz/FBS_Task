package server

import (
	"fmt"
	"github.com/Vitokz/Task/Rest/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (r *Rest) Fibbonaci(c echo.Context) error {
     ctx := c.Request().Context()

     from := c.QueryParam("from")
     if from == "" {
     	from = "0"
	 }

     to := c.QueryParam("to")
     if to == ""{
     	err:=errors.New("Query param \"to\" is empty")
     	 r.Handler.Log.Error(err)
		 return c.JSON(http.StatusBadRequest,models.Err{
		 	Error: fmt.Sprintf("%v",err),
		 })
	 }

     r.Handler.Log.WithFields(logrus.Fields{
     	"event" : "Calculate Fibonacci",
     	"from" : from,
     	"to" : to,
	 }).Info()

     resp,err := r.Handler.Fibonacci(from,to,ctx)
     if err != nil {
		 r.Handler.Log.Error(err)
     	return c.JSON(http.StatusBadRequest,models.Err{
     		Error: fmt.Sprintf("%v",err),
		})
	 }

     return c.JSON(http.StatusOK,resp)
}