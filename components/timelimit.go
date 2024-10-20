package components

import (
    "github.com/yohamta/donburi"
    "typespeed/engine"
    )

type TimeLimitData struct{
    Time *engine.Timeer
}


var Timelimit = donburi.NewComponentType[TimeLimitData]()
