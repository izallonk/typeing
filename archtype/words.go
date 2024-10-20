package archtype

import (
    "github.com/yohamta/donburi"
    "github.com/yohamta/donburi/features/math"
    "github.com/yohamta/donburi/features/transform"
    
    "typespeed/components"
    "typespeed/asset"
    
    )

func NewWords(world donburi.World , pos math.Vec2){
    words := world.Entry(world.Create(
        components.Text,
        transform.Transform,
        components.Border,
        
    ))
    
    components.Text.Get(words).Word = asset.GetRandomWord()
    
    transform.Transform.Get(words).LocalPosition = pos
    word := components.Text.Get(words).Word
    components.Border.SetValue(words, components.BorderData{
        Width: float64 (len(word)* 16),
        Height: 30.0,
    })
    
    
    
}

func NewText(world donburi.World,text string,pos math.Vec2){
     words := world.Entry(world.Create(
        components.Text,
        transform.Transform,
    ))
    components.Text.Get(words).Word = text
    transform.Transform.Get(words).LocalPosition = pos
}

func NewSpawnWord(world donburi.World, pos math.Vec2,spawnfunc components.Spawnfunc){
    spawn := world.Entry(world.Create(components.Spawn,transform.Transform,))

    components.Spawn.Get(spawn).Spawnfunc = spawnfunc
    transform.Transform.Get(spawn).LocalPosition = pos
}



