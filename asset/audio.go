package asset

import (
	"bytes"
	
	"io"
	"os"
	"log"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const sampleRate = 48000

var (
	audioContext *audio.Context
	WordMatch  = map[int]*audio.Player{}
)



func PlaySFX(){
	WordMatch[0].Rewind()
	WordMatch[0].Play()
}

func LoadAudio(){
	audioContext = audio.NewContext(sampleRate)

	WordMatch[0] = Loadwav(audioContext, "./asset/audio/Retro PickUp Coin 04.wav")
	
}

func Loadwav(context *audio.Context, filename string) *audio.Player {
	content,err := os.ReadFile(filename)
    if err !=nil{
        log.Panic ("Fail To Load Audio %s",err)
    }
	decode,err := wav.Decode(context, bytes.NewReader(content))
	if err!= nil{
		log.Panic("Fail To Decode Audio")
	}
	b,err := io.ReadAll(decode)
	if err != nil{
		log.Panic("Fail To Read Audio")
	}
	player := audio.NewPlayerFromBytes(audioContext,b)
	return player
}
