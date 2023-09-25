package gameEngine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func text(engine *EngineStruct) {
	engine.framCount++

	if engine.framCount % 100 == 1 { engine.textBox.frameCount++ }

	if rl.IsKeyPressed(rl.KeySpace) {
		if engine.textBox.textWriting {
			fmt.Printf("1")
		}
		fmt.Printf("2")
	}
}