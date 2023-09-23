package gameEngine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type EngineStruct struct {
	run bool
	bgColor rl.Color

	framCount int

	musicPaused bool
	music rl.Music

	cam rl.Camera2D
	menuSelector bool

	battle bool
	monsterBattle int

	shop bool

	playerTurn bool

	fontText rl.Font
	fontNum rl.Font

	player playerStruct
	bord mapStruct
	sprite spriteStruct
	monster []monsterStruct
	character charcacterStruct
	shopKeeper shopStruct
	menu menuStruct

}

func (engine *EngineStruct) Play() {
	initt(engine)
	for engine.run && !(rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth/2 - engine.sprite.buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + engine.sprite.buttonMenu.Width/2 && screenHeight/2 - engine.sprite.buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + engine.sprite.buttonMenu.Height/2)) {
		menu(engine)
	}

	for engine.run && engine.menuSelector {
		classSelector(engine)
	}

	for engine.run {
		if engine.battle {
			battle(engine)
		} else {
			inMap(engine)
		}
	}
	quit(engine)
}