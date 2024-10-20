package components

import (
    "github.com/yohamta/donburi"
)


type TypeingData struct{
    Enable bool
    
}

var Typeing = donburi.NewComponentType[TypeingData]()
