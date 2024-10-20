package systems

import (
    "github.com/yohamta/donburi"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
   "github.com/hajimehoshi/ebiten/v2"
   "github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi/features/transform"
	"typespeed/components"
	"typespeed/engine"

	"os"
)


type Click struct{
    query *query.Query
    nextscene func()
}

func NewClick(nextscene func())*Click{
    return &Click {
        query : query.NewQuery(filter.Contains(
            components.Border,
    )),
    nextscene : nextscene,
    }
}

func (c *Click) Update(world donburi.World){
    c.query.Each(world, func (entry *donburi.Entry){
        
        border := components.Border.Get(entry)
        pos := transform.WorldPosition(entry)
       
        x := pos.X
        y := pos.Y

        var mouseX,mouseY float64 
        if inpututil. IsMouseButtonJustPressed(ebiten.MouseButton0){
            mouseposX,mouseposY := ebiten.CursorPosition()
            mouseX = float64(mouseposX)
            mouseY = float64(mouseposY)
        
            entryborder := engine.NewRect(x,y,border.Width,border.Height)
            if entryborder.IsInside(mouseX,mouseY){
                if entry.HasComponent(components.Play){
                    c.nextscene()
                }
                if entry.HasComponent(components.Quit){
                    os.Exit(0)
                }
                if entry.HasComponent(components.Retry){
                    c.nextscene()
                }
            }
        }

       
    })
}
