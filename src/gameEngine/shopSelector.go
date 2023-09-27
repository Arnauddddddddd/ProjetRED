package gameEngine

import (
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

func buy(engine *EngineStruct, i int) {
	if len(engine.character.inventory) < 8 && engine.character.gold >= engine.shopKeeper.items[i].price {
		engine.character.inventory = append(engine.character.inventory, engine.shopKeeper.items[i])
		engine.character.gold -= engine.shopKeeper.items[i].price
		if engine.shopKeeper.items[i].gender == "Improvement" {
			engine.shopKeeper.items[i].price *= 2
		}
		if !engine.shopKeeper.items[i].infinitySale {
			engine.shopKeeper.items = append(engine.shopKeeper.items[:i], engine.shopKeeper.items[i+1:]...)
		}
	} else {
		rl.BeginDrawing()
		rl.DrawTextEx(engine.fontText, "Vous Ne Pouvez Pas Acheter Cet Item", rl.NewVector2(float32(engine.player.Dest.X), float32(engine.player.Dest.Y)), 30, 0, rl.Black)
		rl.EndDrawing()
		time.Sleep(2 * time.Second)
	}
}

func shopSelector(engine *EngineStruct) {
	line1 := false
	line2 := false
	line3 := false
	for i := 0; i < len(engine.shopKeeper.items); i++ {
		if i < 3 {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && int32(670 + 89*i) < rl.GetMouseX() && rl.GetMouseX() < int32(670 + 89*i +89) && int32(20+150) < rl.GetMouseY() && rl.GetMouseY() < int32(100+150) {
				buy(engine, i)
			} else if int32(670 + 89*i) < rl.GetMouseX() && rl.GetMouseX() < int32(670 + 89*i +89) && int32(20+150) < rl.GetMouseY() && rl.GetMouseY() < int32(100+150) {
				engine.shopKeeper.showPrice[0] = 1
				engine.shopKeeper.showPrice[1] = i
				line1 = true
			}
		}
	}
	if len(engine.shopKeeper.items) > 3 {
		for j := 3; j < len(engine.shopKeeper.items); j++ {
			if j < 6 {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && int32(670 + 89*(j-3)) < rl.GetMouseX() && rl.GetMouseX() < int32(670 + 89*(j-3)+89) && int32(100+150) < rl.GetMouseY() && rl.GetMouseY() < int32(189+150) {
					buy(engine, j)
				} else if int32(670 + 89*(j-3)) < rl.GetMouseX() && rl.GetMouseX() < int32(670 + 89*(j-3)+89) && int32(100+150) < rl.GetMouseY() && rl.GetMouseY() < int32(189+150) {
					engine.shopKeeper.showPrice[0] = 1
					engine.shopKeeper.showPrice[1] = j
					line2 = true
				}
			}
		}
	}
	if len(engine.shopKeeper.items) > 6 {
		for k := 6; k < len(engine.shopKeeper.items); k++ {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && int32(670 + 89*(k-6)) < rl.GetMouseX() && rl.GetMouseX() < int32(670 + 89*(k-6) +89) && int32(189+150) < rl.GetMouseY() && rl.GetMouseY() < int32(260+160) {
				buy(engine, k)
			} else if int32(670 + 89*(k-6)) < rl.GetMouseX() && rl.GetMouseX() < int32(670 + 89*(k-6) +89) && int32(189+150) < rl.GetMouseY() && rl.GetMouseY() < int32(260+160) {
				engine.shopKeeper.showPrice[0] = 1
				engine.shopKeeper.showPrice[1] = k
				line3 = true
			}
		}
	}
	if !line1 && !line2 && !line3 {
		engine.shopKeeper.showPrice[0] = 0
	}
}
