package asset

import (
    "os"
    "encoding/json"
    "math/rand/v2"
    "log"
    "bytes"
    "github.com/hajimehoshi/ebiten/v2/text/v2"

    
)
type Words struct{
    Words []string
}

var Listofword Words

var Pixboblite *text.GoTextFaceSource

func LoadJson() Words{
    content,err := os.ReadFile("./asset/words/words.json")
    if err !=nil{
        log.Panic ("Fail To Load json %s",err)
    }
    var words Words
    err = json.Unmarshal(content,&words)
    if err != nil{
        log.Panic ("Fail To Unmarhsal json %s",err)
    }
    return words
}

func Loadasset(){
    Listofword = LoadJson()
    Pixboblite = LoadFont()
}


func GetRandomWord() string{
    length := len(Listofword.Words)
    randomindex := rand.IntN(length - 1)
    return Listofword.Words[randomindex]
}


func LoadFont() *text.GoTextFaceSource{
    content ,err := os.ReadFile("./asset/font/ArcticonsSans-Regular.ttf")
    if err != nil{
        log.Panic("Falid to load font file")
    }
    font , err := text.NewGoTextFaceSource(bytes.NewReader(content))
    if err != nil {
        log.Panic("Falid to load font")
    }
    return font
}
