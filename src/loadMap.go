package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func loadMap( mapPath string ) {
	file, err := ioutil.ReadFile(mapPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	remNewLines := strings.Replace(string(file), "\n", " ", -1)
	sliced := strings.Split(remNewLines, " ")
	bord.mapW = -1
	bord.mapH = -1
	for i := 0; i < len(sliced); i++ {
		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		m := int(s)
		if bord.mapW == -1 {
			bord.mapW = m
		} else if bord.mapH == -1 {
			bord.mapH = m
		} else if i < bord.mapW * bord.mapH+2 {
			bord.tileMap = append(bord.tileMap, m)
		} else {
			bord.srcMap = append(bord.srcMap, sliced[i])
		}
	}
	if len(bord.tileMap) > bord.mapW*bord.mapH { bord.tileMap = bord.tileMap[:len(bord.tileMap)-1]}

}