package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "typespeed/scenes"
    "os"
    "log"
    "typespeed/asset"

)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Scene interface{
  Update()
  Draw(screen *ebiten.Image)
}
type Game struct{
    scenes Scene
   
    
}

func (g *Game) Update() error{
   if inpututil.IsKeyJustPressed(ebiten.KeyEscape){
      os.Exit(0)
   }
   if inpututil.IsKeyJustPressed(ebiten.KeySpace){
      g.game()
   }
    g.scenes.Update()
    return nil
}

func (g *Game)Draw(screen *ebiten.Image){

    g.scenes.Draw(screen)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game)mainmenuScene(){
  g.scenes = scenes.NewMainMenu(g.game)
}

func (g *Game) game(){
  
  g.scenes = scenes.NewGame(g.gameover)
}


func (g *Game) gameover(){
  g.scenes = scenes.NewGameover(g.game)
}


func NewGame() *Game{
  asset.Loadasset()   
  asset.LoadAudio()
  game := &Game{}
  game.mainmenuScene()
  return game
}


func main(){
   ebiten.SetWindowSize(800,600)
   ebiten.SetWindowTitle("TypeingTraning")
   if err := ebiten.RunGame(NewGame()); err != nil{
    log.Fatal("Game ERror")
   }
}


