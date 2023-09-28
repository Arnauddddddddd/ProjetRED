package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *EngineStruct) text() { // affiche une zone de text avec du texte à l'interieur
	engine.framCount++

	if engine.framCount%100 == 1 {
		engine.textBox.frameCountSpace++
	} // frame de la barre espace en bas à droite de la textBox
	if engine.framCount%3 == 1 {
		engine.textBox.frameCountText++
	} // frame du d'éfilement du texte

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyPressed(rl.KeySpace) { // si on apuie sur clic gauche ou espace
		if engine.textBox.textWriting { // ET si le texte est entain de s'écrire
			engine.textBox.textPrint = engine.textBox.textToPrint // on affiche le texte
		} else {
			engine.character.showText = false // si le texte est déjà écrie on arrete l'état d'écriture
		}
	}
}
