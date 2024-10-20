package archtype

import (
    "github.com/yohamta/donburi"
    "github.com/yohamta/donburi/features/transform"
    "github.com/yohamta/donburi/features/math"
    "typespeed/components"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"log"
	 "strconv"
)

func NewScore(world donburi.World, pos math.Vec2){
    score := world.Entry(world.Create(
        components.Score,
        components.Text,
        components.Border,
        transform.Transform,
        
    ))

    components.Score.SetValue(score,components.ScoreData{
        Value : 0,
    })
    scorevalue := components.Score.Get(score).Value

    components.Text.SetValue(score,components.TextData{
        Word : strconv.Itoa(scorevalue),
    })
    word := components.Text.Get(score).Word 
    
    transform.Transform.Get(score).LocalPosition = pos
     components.Border.SetValue(score, components.BorderData{
        Width: float64 (len(word)* 16),
        Height: 16.0,
    })
}


func FindScore(world donburi.World) *donburi.Entry{
    score,ok := query.NewQuery(filter.Contains(components.Score,)).First(world)
    if !ok {
        log.Panic("Score Not Found")
    }
    return score
}
