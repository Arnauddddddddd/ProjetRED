package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EngineStruct struct { // création de la structure game engine dans laquels ce déroulera tout le jeu
	run bool
	bgColor rl.Color

	framCount int

	musicPaused bool
	music rl.Music

	cam rl.Camera2D
	menuSelector bool
	menuPlay bool

	doorOpen bool
	doorOpenKey bool

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
	textBox textBoxStruct
	battle battleStruct

}

func (engine *EngineStruct) Play() { // fonction principale
	
	engine.init() // initialisation

	for engine.run && engine.menuPlay { // affichage du menu tant qu'on a pas appuyer sur le bouton Play
		engine.menuFunc()
	}

	for engine.run && engine.menuSelector { // affichage du menu de classe tant qu'on a pas appuyer sur un bouton de classe
		engine.classSelector()
	}

	for engine.run { // boucle principale
		if engine.battle.inBattle { // si on est en combat on affiche la scene de combat sinon on affiche la scene dans la map
			engine.battleFunc()
		} else {
			engine.inMap()
		}
	}
	engine.quit() // unload toute les textures et fichiers ouvert
}