package components

import ("github.com/yohamta/donburi")

type Spawnfunc func (world donburi.World)


type SpawnData struct{
    Spawnfunc Spawnfunc
}

var Spawn = donburi.NewComponentType[SpawnData]()

