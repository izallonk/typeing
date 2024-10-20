package systems

import (
    "github.com/yohamta/donburi"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"typespeed/components"
	//"fmt"
	"image/color"
	//"typespeed/engine"
	//"time"
	
)

type Timelimit struct{
    query *query.Query
   nextscene func()
}

func NewTimelimit(nextscene func()) *Timelimit{
   
    return &Timelimit{
        query : query.NewQuery(filter.Contains(components.Timelimit,)),
        nextscene : nextscene,
    }
}

func (t *Timelimit)Update(world donburi.World){
  
    t.query.Each(world, func (entry *donburi.Entry){
        time := components.Timelimit.Get(entry)
        time.Time.Update()

        if time.Time.IsReady(){
            //Dead Scene
            //time.Time.Restart()
            t.nextscene()
        }
    
    })
}

func (t *Timelimit) Draw(world donburi.World, screen *ebiten.Image){
   
    t.query.Each(world, func (entry *donburi.Entry){
        time := components.Timelimit.Get(entry)
        
        ebitenutil.DrawRect(screen, 20,500,700 * time.Time.PercentDone() ,32,color.RGBA{R:255,G:255,B:255,A:255})
    })
}
