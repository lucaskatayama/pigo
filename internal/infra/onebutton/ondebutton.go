package onebutton

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

const (
	_debounce   = 50 * time.Millisecond
	_clickTicks = 400 * time.Millisecond
	_pressTicks = 800 * time.Millisecond
)

type PushButton struct {
	name       string
	_state     int
	_startTime time.Time

	Pin           gpio.PinIO
	OnLongPress   func()
	OnClick       func()
	OnDoubleClick func()
}

func New(pin string) (*PushButton, error) {
	// Lookup a pin by its number:
	p := gpioreg.ByName(pin)
	if p == nil {
		log.Fatal("Failed to find GPIO2")
		return nil, fmt.Errorf("getting pin [%s]", pin)
	}

	if err := p.In(gpio.PullDown, gpio.BothEdges); err != nil {
		return nil, err
	}

	pb := &PushButton{
		name: pin,
		Pin:  p,
	}
	go func() {
		for {
			pb.Pin.WaitForEdge(_debounce)
			buttonLevel := pb.Pin.Read()
			now := time.Now()
			if pb._state == 0 { // waiting for One pin being pressed.
				if buttonLevel == gpio.High {
					pb._state = 1       // step to state 1
					pb._startTime = now // remember starting time
				}
			} else if pb._state == 1 { // waiting for One pin being released.
				if buttonLevel == gpio.Low {
					pb._state = 2 // step to state 2

				} else if (buttonLevel == gpio.High) && (now.Sub(pb._startTime) > _pressTicks) {
					if pb.OnLongPress != nil {
						pb.OnLongPress()
					}
					pb._state = 6 // step to state 6
				} // if
			} else if pb._state == 2 { // waiting for One pin being pressed the second time or timeout.
				if now.Sub(pb._startTime) > _clickTicks {
					// this was only a single short click
					if pb.OnClick != nil {
						pb.OnClick()
					}
					pb._state = 0 // restart.

				} else if buttonLevel == gpio.High {
					pb._state = 3 // step to state 3
				} // if
			} else if pb._state == 3 { // waiting for One pin being released finally.
				if buttonLevel == gpio.Low {
					// this was a 2 click sequence.
					if pb.OnDoubleClick != nil {
						pb.OnDoubleClick()
					}
					pb._state = 0 // restart.
				} // if
			} else if pb._state == 6 { // waiting for One pin being release after long press.
				if buttonLevel == gpio.Low {
					pb._state = 0 // restart.
				} // if
			}
		}
	}()

	return pb, nil
}

func (pb *PushButton) SetLongPress(h func()) {
	pb.OnLongPress = h
}

func (pb *PushButton) SetDoubleClick(h func()) {
	pb.OnDoubleClick = h
}

func (pb *PushButton) SetClick(h func()) {
	pb.OnClick = h
}
