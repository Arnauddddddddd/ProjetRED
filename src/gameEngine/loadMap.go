package gameEngine

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func loadMap(engine *EngineStruct, mapPath string ) {
	file, err := ioutil.ReadFile(mapPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	remNewLines := strings.Replace(string(file), "\n", " ", -1)
	sliced := strings.Split(remNewLines, " ")
	engine.bord.mapW = -1
	engine.bord.mapH = -1
	for i := 0; i < len(sliced); i++ {
		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		m := int(s)
		if engine.bord.mapW == -1 {
			engine.bord.mapW = m
		} else if engine.bord.mapH == -1 {
			engine.bord.mapH = m
		} else if i < engine.bord.mapW * engine.bord.mapH+2 {
			engine.bord.tileMap = append(engine.bord.tileMap, m+1)
		} else {
			engine.bord.srcMap = append(engine.bord.srcMap, sliced[i])
		}
	}
	if len(engine.bord.tileMap) > engine.bord.mapW*engine.bord.mapH { engine.bord.tileMap = engine.bord.tileMap[:len(engine.bord.tileMap)-1]}

}
