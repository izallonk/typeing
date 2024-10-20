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


type Gameover struct{
    world donburi.World
    systems []System
    drawbles []Drawable
    nextscene func()
}

func NewGameover(nextscene func()) *Gameover{
     render := systems.NewRender()
    debug := systems.NewDebug()


    return &Gameover{
        world : createworldgameover(),
        systems : []System{
          systems.NewClick(nextscene),
          debug,
        },
        drawbles : []Drawable{
            render,
        },
        nextscene : nextscene,
    }
}

func createworldgameover() donburi.World{
    world := donburi.NewWorld()
    archtype.NewText(world, "Game Over", math.NewVec2(320,250))
    archtype.NewRetryButton(world, "Retry", math.NewVec2(360,300))
    archtype.NewQuitButton(world, "Quit", math.NewVec2(360,350))
    world.Entry(world.Create(components.Debug))
    return world
}

func (g *Gameover)Update() {
    for _,s := range g.systems{
        s.Update(g.world)
    }
}

func (g *Gameover)Draw(screen *ebiten.Image){
    for _,d := range g.drawbles{
        d.Draw(g.world,screen)
    }
}

