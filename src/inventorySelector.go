package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
)

func inventorySelector() {
	for i := 0; i < 10; i++ {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) && (int32(497 + 67*i) < rl.GetMouseX() && rl.GetMouseX() < int32(497 + 67*i)+67) && 450 < rl.GetMouseY() && 
		rl.GetMouseY() < int32(520) {
			fmt.Println(i)
		}
	}
}
