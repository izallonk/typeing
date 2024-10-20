package archtype

import (
    "github.com/yohamta/donburi"
    "github.com/yohamta/donburi/features/math"
    "github.com/yohamta/donburi/features/transform"
    
    "typespeed/components"
    //"typespeed/asset"
)


func NewPlayButton(world donburi.World, label string , pos math.Vec2){
    button  := world.Entry(world.Create(
        components.Text,
        components.Border,
        transform.Transform,
        components.Play,
    ))
    
    components.Text.Get(button).Word = label
    transform.Transform.Get(button).LocalPosition = pos
    word := components.Text.Get(button).Word
    components.Border.SetValue(button, components.BorderData{
        Width : float64 ( len(word) * 16),
        Height: 24.0,
    })

}
func NewQuitButton(world donburi.World, label string , pos math.Vec2){
    button  := world.Entry(world.Create(
        components.Text,
        components.Border,
        transform.Transform,
        components.Quit,
    ))
    
    components.Text.Get(button).Word = label
    transform.Transform.Get(button).LocalPosition = pos
    word := components.Text.Get(button).Word
    components.Border.SetValue(button, components.BorderData{
        Width : float64 ( len(word) * 16),
        Height: 24.0,
    })

}

func NewRetryButton(world donburi.World, label string , pos math.Vec2){
    button  := world.Entry(world.Create(
        components.Text,
        components.Border,
        transform.Transform,
        components.Retry,
    ))
    
    components.Text.Get(button).Word = label
    transform.Transform.Get(button).LocalPosition = pos
    word := components.Text.Get(button).Word
    components.Border.SetValue(button, components.BorderData{
        Width : float64 ( len(word) * 16),
        Height: 24.0,
    })

}
