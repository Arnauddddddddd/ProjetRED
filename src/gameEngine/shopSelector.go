package gameEngine

import (
	"fmt"
	
	"github.com/gen2brain/raylib-go/raylib"
)

func shopSelector(engine *EngineStruct) {
	for i := 0; i < len(engine.shopKeeper.items); i++ {
		if i < 3 {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && int32(690 + 80*i) < rl.GetMouseX() && rl.GetMouseX() < int32(690 + 80*i +80) && int32(20) < rl.GetMouseY() && rl.GetMouseY() < int32(100) {
				fmt.Println(i)
			}
		}
	}
	if len(engine.shopKeeper.items) > 3 {
		for j := 3; j < len(engine.shopKeeper.items); j++ {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && int32(690 + 80*(j-3)) < rl.GetMouseX() && rl.GetMouseX() < int32(690 + 80*(j-3)+80) && int32(100) < rl.GetMouseY() && rl.GetMouseY() < int32(180) {
				fmt.Println(j)
			}
		}
	}
	if len(engine.shopKeeper.items) > 6 {
		for k := 6; k < len(engine.shopKeeper.items); k++ {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && int32(690 + 80*(k-6)) < rl.GetMouseX() && rl.GetMouseX() < int32(690 + 80*(k-6) +80) && int32(180) < rl.GetMouseY() && rl.GetMouseY() < int32(260) {
				fmt.Println(k)
			}
		}
	}
}