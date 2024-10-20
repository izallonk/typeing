package components

import ("github.com/yohamta/donburi")


type TextData struct{
    Word string
    Runes []rune
}

var Text = donburi.NewComponentType[TextData]()
