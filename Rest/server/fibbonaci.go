package server

import (
	"fmt"
	"github.com/Vitokz/Task/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (r *Rest) Fibbonaci(c echo.Context) error {

	from := c.QueryParam("from")
	if from == "" {
		from = "1"
	}
	fromInt, err := strconv.Atoi(from)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Err{
			Error: fmt.Sprintf("%v", errors.New("\"from\" param is not number")),
		})
	} else if fromInt < 1 {
		return c.JSON(http.StatusBadRequest, models.Err{
			Error: fmt.Sprintf("%v", errors.New("\"from\" param is minus or null")),
		})
	}

	to := c.QueryParam("to")
	if to == "" {
		err := errors.New("Query param \"to\" is empty")
		r.Handler.Log.Error(err)
		return c.JSON(http.StatusBadRequest, models.Err{
			Error: fmt.Sprintf("%v", err),
		})
	}
	toInt, err := strconv.Atoi(to)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Err{
			Error: fmt.Sprintf("%v", errors.New("\"to\" param is not number")),
		})
	} else if toInt < 0 {
		return c.JSON(http.StatusBadRequest, models.Err{
			Error: fmt.Sprintf("%v", errors.New("\"to\" param is minus")),
		})
	} else if toInt < fromInt {
		return c.JSON(http.StatusBadRequest, models.Err{
			Error: fmt.Sprintf("%v", errors.New("\"to\" param is less than \"from\"")),
		})
	} else if toInt > 92 {
		return c.JSON(http.StatusBadRequest, models.Err{
			Error: fmt.Sprintf("%v", errors.New("\"to\" param must be less than or equal to 92")),
		})
	}

	r.Handler.Log.WithFields(logrus.Fields{
		"event": "Calculate Fibonacci",
		"from":  from,
		"to":    to,
	}).Info()

	resp, err := r.Handler.Fibonacci(fromInt, toInt)
	if err != nil {
		r.Handler.Log.Error(err)
		return c.JSON(http.StatusBadRequest, models.Err{
			Error: fmt.Sprintf("%v", err),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
