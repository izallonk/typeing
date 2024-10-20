package events

import (
    "github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
	"typespeed/archtype"
	"typespeed/asset"
	"typespeed/components"
	//"strings"
	"strconv"
)


type Wordmatch struct{
    Entry *donburi.Entry
}

var WordmatchEvent = events.NewEventType[Wordmatch]()


func OnWordMatch(world donburi.World, event Wordmatch){
		
		score := archtype.FindScore(world)
		scorevalue := components.Score.Get(score)
		scoretext := components.Text.Get(score)

		timelimit := archtype.FindTimelimit(world)
		time := components.Timelimit.Get(timelimit)

		scorevalue.Value ++
		scoretext.Word =  strconv.Itoa(scorevalue.Value)
		time.Time.Restart()
		asset.PlaySFX()
		event.Entry.Remove()
}

func SetupEvent(world donburi.World){
		WordmatchEvent.Subscribe(world, OnWordMatch)
}

type Events struct{}

func NewEvents() *Events{
	return &Events{}
}

func (e *Events)Update(world donburi.World){
	WordmatchEvent.ProcessEvents(world)
}
