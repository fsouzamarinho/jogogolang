package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
	"math/rand"
	"time"
	"fmt"
)

const (
	screenWidth   = 640
	screenHeight  = 480
	playerWidth   = 64
	playerHeight  = 64
	playerSpeed   = 4
	enemySize     = 40
	spawnInterval = 30 // frames
)

type Enemy struct {
	X, Y   float64
	SpeedY float64
}

type Game struct {
	playerX, playerY float64
	playerImg        *ebiten.Image
	enemies          []Enemy
	frameCount       int
	gameOver         bool
	startTime        time.Time
	score            float64
}

func (g *Game) restart() {
	g.playerX = screenWidth / 2
	g.playerY = screenHeight - playerHeight - 10
	g.enemies = []Enemy{}
	g.startTime = time.Now()
	g.score = 0
	g.gameOver = false
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func (g *Game) Update() error {

	if g.gameOver {
    if ebiten.IsKeyPressed(ebiten.KeyR) {
        g.restart()
    }
    return nil
}
	g.score = time.Since(g.startTime).Seconds()

	// Movimento do jogador
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.playerX > 0 {
		g.playerX -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.playerX < screenWidth-playerWidth {
		g.playerX += playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.playerY > 0 {
		g.playerY -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.playerY < screenHeight-playerHeight {
		g.playerY += playerSpeed
	}

	// Incrementa o contador de frames
	g.frameCount++

	// Gera inimigo novo a cada intervalo
	if g.frameCount%spawnInterval == 0 {
		newEnemy := Enemy{
			X:      float64(rand.Intn(screenWidth - enemySize)),
			Y:      -float64(enemySize),
			SpeedY: 2 + rand.Float64()*3,
		}
		g.enemies = append(g.enemies, newEnemy)
	}

	// Atualiza posição dos inimigos
	for i := range g.enemies {
		g.enemies[i].Y += g.enemies[i].SpeedY
	}

	// Remove inimigos que saíram da tela
	activeEnemies := g.enemies[:0]
	for _, e := range g.enemies {
		if e.Y < screenHeight {
			activeEnemies = append(activeEnemies, e)
		}
	}
	g.enemies = activeEnemies

	for _, e := range g.enemies {
		if checkCollision(g.playerX, g.playerY, e.X, e.Y) {
			g.gameOver = true
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(image.Black)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Tempo: %.1fs", g.score))

	// Desenha jogador
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerX, g.playerY)
	screen.DrawImage(g.playerImg, op)

	// Desenha inimigos
	for _, e := range g.enemies {
		ebitenutil.DrawRect(screen, e.X, e.Y, enemySize, enemySize, image.White)
	}

	if g.gameOver {
		 ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Pontuação final: %.1fs", g.score), screenWidth/2-70, screenHeight/2+20)
    ebitenutil.DebugPrintAt(screen, "Pressione R para reiniciar", screenWidth/2-90, screenHeight/2+40)
		ebitenutil.DebugPrintAt(screen, "GAME OVER", screenWidth/2-40, screenHeight/2)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func checkCollision(px, py float64, ex, ey float64) bool {
	return px < ex+enemySize &&
		px+playerWidth > ex &&
		py < ey+enemySize &&
		py+playerHeight > ey
}

func main() {

	rand.Seed(time.Now().UnixNano())

	game := &Game{
		playerX:    screenWidth / 2,
		playerY:    screenHeight - playerHeight - 10,
		playerImg:  loadImage("player.png"),
		enemies:    []Enemy{},
		frameCount: 0,
		startTime:  time.Now(),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Etapa 2 - Inimigos Caindo")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
