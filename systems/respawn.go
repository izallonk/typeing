package systems

import (
    "github.com/yohamta/donburi"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"github.com/yohamta/donburi/features/math"
	"typespeed/components"
	"typespeed/archtype"
	"typespeed/engine"

)

type Respawn struct {
	query *query.Query
}

func NewRespawn() *Respawn{
	return &Respawn{
		query : query.NewQuery(filter.Contains(components.Text,components.Border)),
	}
}

func (r *Respawn) Update(world donburi.World){
	var words []*donburi.Entry
	r.query.Each(world, func (entry *donburi.Entry){
		if entry.HasComponent(components.Score){
			return
		}
		words = append(words,entry)
		
	})

	if len(words) <=0{
		
		for range  10{
			 randompos := math.NewVec2(engine.Randomrange(40,400),engine.Randomrange(40,400))
			 archtype.NewSpawnWord(world,randompos,SpawnWords(randompos))
		}
	}
}

func SpawnWords(randompos math.Vec2) func (world donburi.World){
   
   return func (world donburi.World){
      archtype.NewWords(world, randompos)
   }
}
