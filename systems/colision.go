package systems

import (
	"github.com/yohamta/donburi"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	
	"github.com/yohamta/donburi/features/transform"
	"typespeed/components"
	"typespeed/engine"

)

type Colision struct{
	query *query.Query
}

func Newcolision() *Colision{
	return &Colision{
		query : query.NewQuery(filter.Contains(components.Border,transform.Transform,)),
	}
}

func (c *Colision) Update(world donburi.World){
	var entrys []*donburi.Entry
	
	c.query.Each(world, func (entry *donburi.Entry){
		entrys = append(entrys,entry)
	})
	

	for _,entry := range entrys{
		if !entry.Valid(){
			continue
		}
		entryborder := components.Border.Get(entry)
		for _,other := range entrys{
			if entry.Entity().Id() == other.Entity().Id(){
				continue
			}
			if !entry.Valid() || !other.Valid(){
			continue
		}
			otherborder := components.Border.Get(other)

			posentry := transform.Transform.Get(entry).LocalPosition
			posother := transform.Transform.Get(other).LocalPosition
			entryRect := engine.NewRect(posentry.X,posentry.Y,entryborder.Width,entryborder.Height)
			otherRect := engine.NewRect(posother.X,posother.Y,otherborder.Width,otherborder.Height)

			if entryRect.Colision(otherRect){
				entry.Remove()
			}
		}
	}
}
