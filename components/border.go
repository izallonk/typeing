package components

import ("github.com/yohamta/donburi")

type BorderData struct{
    Width float64
    Height float64
    
}

var Border = donburi.NewComponentType[BorderData]()
