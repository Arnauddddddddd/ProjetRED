package gameEngine

import 	rl "github.com/gen2brain/raylib-go/raylib"

func (engine *EngineStruct) classSelector() {
	engine.run = !rl.WindowShouldClose()
	rl.BeginDrawing()

    engine.framCount++

    if engine.framCount % 10 == 1 { engine.menu.frameCount++ } // acctualise le frame Count du background
	if engine.menu.frameCount > 19 { engine.menu.frameCount = 0 } // si il est trop grand on le réinitialise à 0 (la 1er frame)
	engine.menu.Src.X = engine.menu.Src.Width * float32(engine.menu.frameCount) // change la source (engine.menu.Src) de l'image du bg afficher

	rl.DrawTexturePro(engine.menu.sprite, engine.menu.Src, engine.menu.Dest, rl.NewVector2(engine.menu.Dest.Width, engine.menu.Dest.Height), 0, rl.White) // affiche le bg

    // affiche les nom des classes
    rl.DrawTextEx(engine.fontText, "Thief", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+100)), 60, 0, rl.Black) 
    rl.DrawTextEx(engine.fontText, "Slayer", rl.NewVector2(float32(screenWidth/2 - engine.sprite.buttonPlay.Width/2-30), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+100)), 60, 0, rl.Black)
    rl.DrawTextEx(engine.fontText, "Tank", rl.NewVector2(float32(screenWidth - engine.sprite.buttonPlay.Width*2-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+100)), 60, 0, rl.Black)

    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (engine.sprite.buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < engine.sprite.buttonPlay.Width*2 && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 + engine.sprite.buttonPlay.Height) { // initialise le personage de la classe voleur si on appuie sur le boutton
        engine.menuSelector = false
        engine.character.name = "Thief"
        engine.character.hp = 150
        engine.character.damage = 30
        engine.character.damageBase = 30
        engine.character.speed = 10
        engine.character.hpMax = 150
        engine.character.gold = 0
        engine.character.alive = true
    }
    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth/2 - engine.sprite.buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + engine.sprite.buttonPlay.Width/2 && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height) { // initialise le personage de la classe Slayer si on appuie sur le boutton
        engine.menuSelector = false
        engine.character.name = "Slayer"
        engine.character.hp = 150
        engine.character.damage = 40
        engine.character.damageBase = 40
        engine.character.speed = 10
        engine.character.hpMax = 150
        engine.character.gold = 0
        engine.character.alive = true
        engine.character.inventory = append( engine.character.inventory, itemStruct{"Strengt item", "Fork", "By using this fork, you fight \nthe rest of the fight with it and will \ninflict an additional 50 damage on you.", rl.LoadTexture("../texture/shop/fork.png"), 0, 0, 0, false, true, false,1})
        engine.character.inventory = append( engine.character.inventory, itemStruct{"Strengt item", "Fork", "By using this fork, you fight \nthe rest of the fight with it and will \ninflict an additional 50 damage on you.", rl.LoadTexture("../texture/shop/fork.png"), 0, 0, 0, false, true, false,1})
    }
    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth - engine.sprite.buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - engine.sprite.buttonPlay.Width && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height) { // initialise le personage de la classe Tank si on appuie sur le boutton
        engine.menuSelector = false
        engine.character.name = "Tank"
        engine.character.hp = 200
        engine.character.damage = 30
        engine.character.damageBase = 30
        engine.character.speed = 10
        engine.character.hpMax = 250
        engine.character.gold = 0
        engine.character.alive = true
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Heal Potion", "Potion", "A simple magic potion that restores you 50 hp", rl.LoadTexture("../texture/shop/potion.png"), 0, 50, 0, true, true, true, 20})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Heal Potion", "Potion", "A simple magic potion that restores you 50 hp", rl.LoadTexture("../texture/shop/potion.png"), 0, 50, 0, true, true, true, 20})
    }



	if engine.sprite.buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < engine.sprite.buttonPlay.Width*2 && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height { // affiche le boutton enfoncer si on passe notre souris dessus 
        rl.DrawTexture(engine.sprite.buttonPlay, engine.sprite.buttonPlay.Width, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
        rl.DrawTextEx(engine.fontNum, "         HP : 150       Damages : 30     ", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+350)), 30, 0, rl.Black) 
        rl.DrawTextEx(engine.fontNum, "-You start without objects", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+450)), 25, 0, rl.Black) 
        rl.DrawTextEx(engine.fontNum, "-You gain 25% more gold after each fight", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+500)), 25, 0, rl.Black) 
        rl.DrawTextEx(engine.fontNum, "FATALITY : You steal 40% of the monster's life", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+600)), 25, 0, rl.Black) 

    } else {
        rl.DrawTexture(engine.sprite.buttonPlayPressed, engine.sprite.buttonPlay.Width, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
    }

	if screenWidth/2 - engine.sprite.buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + engine.sprite.buttonPlay.Width/2 && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height{ // affiche le boutton enfoncer si on passe notre souris dessus 
        rl.DrawTexture(engine.sprite.buttonPlay, screenWidth/2 - engine.sprite.buttonPlay.Width/2, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
        rl.DrawTextEx(engine.fontNum, "         HP : 150       Damages : 40     ", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+350)), 30, 0, rl.Black) 
        rl.DrawTextEx(engine.fontNum, "-You start with two forks which inflict 100 damages", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+450)), 25, 0, rl.Black) 
        rl.DrawTextEx(engine.fontNum, "-You are more powerful (+33%)", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+500)), 25, 0, rl.Black) 
        rl.DrawTextEx(engine.fontNum, "FATALITY : After using, your damages inflict 50% more", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+600)), 25, 0, rl.Black) 

    } else {
        rl.DrawTexture(engine.sprite.buttonPlayPressed, screenWidth/2 - engine.sprite.buttonPlay.Width/2, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
    }

	if screenWidth - engine.sprite.buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - engine.sprite.buttonPlay.Width && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height{ // affiche le boutton enfoncer si on passe notre souris dessus 
        rl.DrawTexture(engine.sprite.buttonPlay, screenWidth - engine.sprite.buttonPlay.Width*2, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
        rl.DrawTextEx(engine.fontNum, "         HP : 200       Damages : 30     ", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+350)), 30, 0, rl.Black) 
        rl.DrawTextEx(engine.fontNum, "-You start with two heal potions which regen you 50 HP ", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+450)), 25, 0, rl.Black) 
        rl.DrawTextEx(engine.fontNum, "-You have more life (+33%)", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+500)), 25, 0, rl.Black) 
        rl.DrawTextEx(engine.fontNum, "FATALITY : After using, you take 40% less damages", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+600)), 25, 0, rl.Black) 

    } else {
        rl.DrawTexture(engine.sprite.buttonPlayPressed, screenWidth - engine.sprite.buttonPlay.Width*2, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
    }





	rl.ClearBackground(engine.bgColor)
	rl.EndDrawing()
}