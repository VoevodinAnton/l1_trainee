package tasks

import (
	"fmt"
	"strconv"
)

type LightBulb struct {
	IsOn bool
}

func (lightBulb *LightBulb) GetIsOn() bool {
	return lightBulb.IsOn
}

func (lightBulb *LightBulb) SetIsOn(isOn bool) {
	lightBulb.IsOn = isOn
}

func (lightBulb *LightBulb) String() string {
	return "LightBulb: " + strconv.FormatBool(lightBulb.IsOn)
}

type Lighting interface {
	Light()
}

type LightSwitchAdapter struct {
	lightBulb LightBulb
}

func (lightSwitchAdapter *LightSwitchAdapter) Light() {
	lightBulb := lightSwitchAdapter.lightBulb

	if lightBulb.GetIsOn() {
		fmt.Println("Light is on, do nothing")
	} else {
		lightBulb.SetIsOn(true)
		fmt.Println("Light is off, turning it on")
	}
}

func main() {
	var lightSwitch Lighting = &LightSwitchAdapter{
		lightBulb: LightBulb{true},
	}

	lightSwitch.Light()

}
