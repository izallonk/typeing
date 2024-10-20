package systems


import (
    "github.com/yohamta/donburi"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"typespeed/components"
	
)

type Spawn struct{
    query *query.Query
}

func NewSpawn() *Spawn{
    return &Spawn{
        query : query.NewQuery(filter.Contains(components.Spawn,)),
    }
}

func (s *Spawn) Update(world donburi.World){
    s.query.Each(world, func (entry *donburi.Entry){
        spawn := components.Spawn.Get(entry)
        spawn.Spawnfunc(world)
        entry.Remove()
    })
}
