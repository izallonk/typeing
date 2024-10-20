package systems

import (
    "github.com/yohamta/donburi"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"github.com/yohamta/donburi/features/transform"
    "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"typespeed/components"
	"log"
	"image/color"
)



type Debug struct{
    query *query.Query
    debug *components.DebugData
}

func NewDebug() *Debug{
    return &Debug{
        query : query.NewQuery(filter.Contains(components.Border,transform.Transform)),
    }
}

func(d *Debug)Update(world donburi.World){
     if d.debug == nil{
        debug , ok := query.NewQuery(filter.Contains(components.Debug)).First(world)
        if !ok{
            log.Panic("Debug Not Found")
        }
        d.debug = components.Debug.Get(debug)
     }
     if inpututil.IsKeyJustPressed(ebiten.KeySlash){
        d.debug.Enable = ! d.debug.Enable
     }
}


func (d *Debug)Draw(world donburi.World, screen *ebiten.Image){
    if d.debug == nil || d.debug.Enable{
        d.query.Each(world, func (entry *donburi.Entry){
            border := components.Border.Get(entry)
            pos := transform.WorldPosition(entry)
            
            ebitenutil.DrawRect(screen,pos.X,pos.Y,border.Width,border.Height,color.RGBA{R:12,G:32,B:44,A:255})
        })
    }
}


