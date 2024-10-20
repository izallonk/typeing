package systems

import (
    "github.com/yohamta/donburi"
    "github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"typespeed/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Typing struct{
    query *query.Query
}

func NewTyping() *Typing{
    return &Typing{
    query : query.NewQuery(filter.Contains(
        components.Text,
        components.Typeing,
    )),}
}

func (t *Typing) Update(world donburi.World){
    t.query.Each(world, func (entry *donburi.Entry){
        text := components.Text.Get(entry)
        
        text.Runes = ebiten.AppendInputChars(text.Runes[:0])
        text.Word += string(text.Runes)
        if repeatingKeyPressed(ebiten.KeyBackspace) {
            if len(text.Word) >= 1 {
                text.Word = text.Word[:len(text.Word)-1]
            }
        }
        
    })
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
