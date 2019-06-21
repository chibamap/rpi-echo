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

func (self *Led) TurnOn(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(fmt.Sprint("unable to refer id", err.Error()))
	}
	pin := self.Pins[id].High()

	return c.String(http.StatusOK, fmt.Sprintf("turned on led %s", id))
}

func (self *Led) TurnOff(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func (self *Led) Close() {
	rpio.Close()
}
