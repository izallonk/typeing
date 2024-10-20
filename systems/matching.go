package systems

import (
	"github.com/yohamta/donburi"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"typespeed/components"
	"typespeed/archtype"
	"typespeed/events"
	"strings"

	
)


type Matching struct{
	query *query.Query
}

func NewMatching() *Matching{
	return &Matching{
		query : query.NewQuery(filter.Contains(
			components.Text,
		)),
	}
}

func (m *Matching) Update(world donburi.World){
	m.query.Each(world, func (entry *donburi.Entry){
		typeing := archtype.FindTypeing(world)
		typeingtext := components.Text.Get(typeing)
		
		
		if entry.Entity().Id() == typeing.Entity().Id(){
			return
		}
		texttype := typeingtext.Word
		if  strings.Contains(texttype,"_"){
			 texttype = texttype[:len(texttype)-1]
		}
		
		wordtext := components.Text.Get(entry)
		
		if texttype == wordtext.Word{
			typeingtext.Word = ""
			events.WordmatchEvent.Publish(world, events.Wordmatch{
				Entry : entry,
			})
		}
	
	})
}
