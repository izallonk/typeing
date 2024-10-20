package engine

import (
    "time"
)

type Timeer struct{
    Curentframe int
    Targetframe int
}

func NewTimer(d time.Duration) *Timeer{
	return &Timeer{
		Curentframe : 0,
		Targetframe : int(d.Milliseconds()) * 60 / 1000,
	}
}

func (t *Timeer) Update(){
	if t.Curentframe < t.Targetframe{
		t.Curentframe++
	}
}

func (t *Timeer)IsReady()bool{
	return t.Curentframe>= t.Targetframe
}

func (t *Timeer)Restart(){
	t.Curentframe = 0
}
func (t *Timeer) PercentDone() float64 {
	return float64(t.Curentframe) / float64(t.Targetframe)
}


