package scenes

import (
    "github.com/hajimehoshi/ebiten/v2"
    //"github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
      "github.com/yohamta/donburi/features/math"
    "github.com/yohamta/donburi"
   "typespeed/archtype"
    "typespeed/systems"
    //"typespeed/asset"
    "typespeed/events"
    "typespeed/components"
    
)

const (
	screenWidth  = 800
	screenHeight = 600
)
type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}

type Scene struct{
    World donburi.World
    systems []System
    drawbles []Drawable
    nextscene func()
    
}

func (s *Scene) Update() {

   if inpututil.IsKeyJustPressed(ebiten.KeySpace){
      s.Restart()
   }
  
    for _,systems := range s.systems{
        systems.Update(s.World)
    }

}

func (s *Scene)Draw(screen *ebiten.Image){
    for _,draw := range s.drawbles{
        draw.Draw(s.World,screen)
    }
}
func NewGame( nextscene func()) *Scene{
 
    world := CreateWorld()
    render := systems.NewRender()
      debug := systems.NewDebug()
    
    return &Scene{
        World: world,
        systems : []System{
            systems.NewTyping(),
            systems.NewSpawn(),
            systems.NewMatching(),
            systems.Newcolision(),
            systems.NewRespawn(),
            systems.NewTimelimit(nextscene),
            events.NewEvents(),  
            debug,
        },
        drawbles : []Drawable{
            systems.NewTimelimit(nextscene),
            debug,
            render,
        },
       nextscene : nextscene,
    }
}

func CreateWorld() donburi.World{
   world := donburi.NewWorld()
   archtype.NewScore(world, math.NewVec2(400,20))
   archtype.NewTypeing(world)
   archtype.NewTimelimit(world)
   events.SetupEvent(world)
   world.Entry(world.Create(components.Debug))
   return world
}


func (s *Scene)Restart(){
   s.World = nil
   s.World = CreateWorld()
}



