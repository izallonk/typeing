package archtype

import (
    "github.com/yohamta/donburi"
    //"github.com/yohamta/donburi/features/transform"
    //"github.com/yohamta/donburi/features/math"
    "typespeed/components"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"time"
	"typespeed/engine"
	"log"
	
)


func NewTimelimit(world donburi.World){
    timelimit := world.Entry(world.Create(
        components.Timelimit,
        
    ))
   var countdown time.Duration = 30 * time.Second
    components.Timelimit.SetValue(timelimit, components.TimeLimitData{
        Time : engine.NewTimer(countdown),
    })
}

func FindTimelimit(world donburi.World) *donburi.Entry{
    timelimit , ok := query.NewQuery(filter.Contains(components.Timelimit)).First(world)
    if !ok {
        log.Panic("Timelimit Not FOund")
    }
    return timelimit
}

