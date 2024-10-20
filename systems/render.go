package systems

import (
    "github.com/yohamta/donburi"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"github.com/yohamta/donburi/features/transform"
	
	"typespeed/components"
    "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
    "typespeed/engine"
    "typespeed/asset"
    "time"

    "golang.org/x/text/language"
    )

type Render struct{
    query *query.Query
    timer *engine.Timeer
}

func NewRender() *Render{
      var time time.Duration= 1 * time.Second
    return &Render{
    query : query.NewQuery(
        filter.Contains(components.Text,transform.Transform)),
    timer : engine.NewTimer(time),
    }
    
    
}

func (r *Render) Draw(world donburi.World, screen *ebiten.Image){
    r.timer.Update()
    r.query.Each(world, func(entry *donburi.Entry){
        texts := components.Text.Get(entry)
        pos := transform.WorldPosition(entry)
        //blink
        word := texts.Word
        if entry.HasComponent(components.Typeing){
            if r.timer.PercentDone() < 0.40{
                word += "_"
               
            }
            
            if r.timer.IsReady(){
                r.timer.Restart()
            }
        }
        font := &text.GoTextFace{
            Source : asset.Pixboblite,
            Size : 24,
            Language: language.English,
        }
        x := pos.X
        y := pos.Y
        op := &text.DrawOptions{}
        op.GeoM.Translate(x,y)
        text.Draw(screen, word,font,op)
    })
}
