package gameEngine

func (engine *EngineStruct) battleFunc() {
	engine.updateBattle() // met à jour les variables
	if engine.player.showHud {
		engine.inventorySelector(640, 705, 795, 795, 890, 87) // check si on appuie dans l'inventaire à notre tour Si le hub est affiché
	}
	engine.drawSceneBattle() // affiche la scene
}