package gameEngine

func battle(engine *EngineStruct) {
	updateBattle(engine)
	inventorySelector(engine, 640, 705, 795, 795, 890, 87)
	drawSceneBattle(engine)
}