package main

import (
	"strconv"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth = 1000
	screenHeight = 480
)

var (
	run = true
	bgColor = rl.NewColor(147, 211, 196, 255)

	grassSprite rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc rl.Rectangle
	playerDest rl.Rectangle

	playerMoving bool
	playerDir int
	playerUp, playerDown, playerRight, playerLeft bool
	playerFrame int

	framCount int

	playerSpeed float32 = 3

	musicPaused bool
	music rl.Music

	cam rl.Camera2D

    buttonMenu rl.Texture2D
    buttonMenuPressed rl.Texture2D

)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
    rl.DrawText(strconv.Itoa(int(rl.GetMouseX())) + " " + strconv.Itoa(int(rl.GetMouseY())), 0, 0, 20, rl.Black)
		
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerMoving = true
		playerDir = 1
		playerUp = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving = true
		playerDir = 0
		playerDown = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerDir = 2
		playerLeft = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerDir = 3
		playerRight = true
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}
}

func update() {
	run = !rl.WindowShouldClose()

	playerSrc.X = playerSrc.Width * float32(playerFrame)

	if playerMoving {
		if playerUp { playerDest.Y -= playerSpeed }
		if playerDown { playerDest.Y += playerSpeed }
		if playerLeft { playerDest.X -= playerSpeed }
		if playerRight { playerDest.X += playerSpeed }
		if framCount % 8 == 1 { playerFrame++ }
	} else if framCount % 45 == 1 {
		playerFrame++
	}

	framCount++
	if playerFrame > 3 { playerFrame = 0 }
	if !playerMoving && playerFrame > 1 { playerFrame = 0 }

	playerSrc.X = playerSrc.Width * float32(playerFrame)
	playerSrc.Y = playerSrc.Height * float32(playerDir)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(playerDest.X - (playerDest.Width / 2)), float32(playerDest.Y - (playerDest.Height/2)))

	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bgColor)

	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Le Jeu")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("res/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("res/Characters/Basic Charakter Spritesheet.png")
    buttonMenu = rl.LoadTexture("button_menu/png/Buttons/Rect-Text-Blue/Play-Idle.png")
    buttonMenuPressed = rl.LoadTexture("button_menu/png/Buttons/Rect-Text-Blue/Play-Click.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)



	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/music.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(playerDest.X - (playerDest.Width / 2)), float32(playerDest.Y - (playerDest.Height/2))), 0.0, 1.0)

}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)

	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	
	rl.CloseWindow()
}

func menu() {
	rl.BeginDrawing()
    if screenWidth/2 - buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + buttonMenu.Width/2 && screenHeight/2 - buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + buttonMenu.Height/2 {
        rl.DrawTexture(buttonMenuPressed, screenWidth/2 - buttonMenuPressed.Width/2, screenHeight/2 - buttonMenu.Height/2, rl.White)
    } else {
        rl.DrawTexture(buttonMenu, screenWidth/2 - buttonMenu.Width/2, screenHeight/2 - buttonMenu.Height/2, rl.White)
    }
    
	rl.ClearBackground(bgColor)
	rl.DrawText("Play ", 400, 20, 20, rl.Black)
    rl.DrawText(strconv.Itoa(int(rl.GetMouseX())) + " " + strconv.Itoa(int(rl.GetMouseY())), 0, 0, 20, rl.Black)

	rl.EndDrawing()
}

func main() {

	for !(rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth/2 - buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + buttonMenu.Width/2 && screenHeight/2 - buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + buttonMenu.Height/2)) {
		menu()
	}

	for run {
		input()
		update()
		render()
	}
	quit()
}