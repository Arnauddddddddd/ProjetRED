package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func text(engine *EngineStruct) {
	engine.framCount++

	if engine.framCount % 100 == 1 { engine.textBox.frameCountSpace++ }
	if engine.framCount % 5 == 1 { engine.textBox.frameCountText++ }

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyPressed(rl.KeySpace){
		if engine.textBox.textWriting {
			engine.textBox.textPrint = engine.textBox.textToPrint
		} else {
			engine.character.showText = false
		}
	}
}