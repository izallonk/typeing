package archtype

import (
    "github.com/yohamta/donburi"
    "github.com/yohamta/donburi/features/transform"
     "github.com/yohamta/donburi/features/math"
    "typespeed/components"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"log"
)


func NewTypeing(world donburi.World){
    typeing := world.Entry(world.Create(
        components.Text,
        transform.Transform,
        components.Typeing,
    ))

    components.Typeing.SetValue(typeing,components.TypeingData{
        Enable : true,
    })

    transform.Transform.Get(typeing).LocalPosition = math.NewVec2(50,550)
}

func FindTypeing(world donburi.World) *donburi.Entry{
    typeing,ok := query.NewQuery(filter.Contains(components.Typeing,)).First(world)
    if !ok{
        log.Panic("Painc Not Found")
    }
    return typeing
    
}



