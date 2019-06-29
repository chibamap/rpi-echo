package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stianeikeland/go-rpio/v4"
	"net/http"
	"strconv"
)

type (
	Led struct {
		Pins map[int]rpio.Pin
	}
)

func New() *Led {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	led := &Led{}
	pinIds := []int{16, 20, 21, 26}
	for id := range pinIds {
		pin := rpio.Pin(id)
		pin.Output()
		led.Pins[id] = pin
	}

	return led
}

func (self *Led) Turn(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(fmt.Sprint("unable to refer id", err.Error()))
	}
	pin := self.Pins[id]
	pin.Toggle()

	return c.String(http.StatusOK, fmt.Sprintf("toggle led %s", id))
}

func (self *Led) Close() {
	rpio.Close()
}
