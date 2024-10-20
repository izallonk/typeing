package components

import (
    "github.com/yohamta/donburi"
)

type ScoreData struct{
    Value int
}

var Score = donburi.NewComponentType[ScoreData]()
