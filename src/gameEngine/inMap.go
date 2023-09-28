package gameEngine

func (engine *EngineStruct) inMap() {
	engine.input() // check les input
	engine.update() // met a jour les variables
	engine.drawScene() // affiche la scene
}