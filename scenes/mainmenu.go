package scenes


import (
    "github.com/hajimehoshi/ebiten/v2"
    //"github.com/hajimehoshi/ebiten/v2/ebitenutil"
    //"github.com/hajimehoshi/ebiten/v2/inpututil"
      "github.com/yohamta/donburi/features/math"
    "github.com/yohamta/donburi"
   "typespeed/archtype"
    "typespeed/systems"
    //"typespeed/asset"
    //"typespeed/events"
    "typespeed/components"
)


type MainMenu struct{

    world donburi.World
    systems []System
    drawbles []Drawable
    nextscene func()
}

func NewMainMenu(nextscene func ()) *MainMenu{
    render := systems.NewRender()
    debug := systems.NewDebug()
    
    return &MainMenu{
        world : createworldmainmenu(),
        systems : []System{
            systems.NewClick(nextscene),
            debug,
            
        },
        drawbles : []Drawable{
            debug,
            render,
        },
        nextscene : nextscene,
    }
}

func createworldmainmenu() donburi.World{
    world := donburi.NewWorld()
    archtype.NewPlayButton(world, "Play", math.NewVec2(360,300))
    archtype.NewQuitButton(world, "Quit", math.NewVec2(360,350))
    world.Entry(world.Create(components.Debug))
    return world
}

func (m *MainMenu)Update() {
    for _,s := range m.systems{
        s.Update(m.world)
    }
}

func (m *MainMenu)Draw(screen *ebiten.Image){
    for _,d := range m.drawbles{
        d.Draw(m.world,screen)
    }
}

