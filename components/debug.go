package components

import ("github.com/yohamta/donburi")


type DebugData struct{
    Enable bool
}

var Debug = donburi.NewComponentType[DebugData]()
