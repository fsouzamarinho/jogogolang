package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
	playerWidth  = 64
	playerHeight = 64
	speed        = 3
)

type Game struct {
	playerX float64
	playerY float64
	player  *ebiten.Image
}

func (g *Game) Update() error {
if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.playerX > 0 {
    g.playerX -= speed
}
if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.playerX < screenWidth - playerWidth {
    g.playerX += speed
}
if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.playerY > 0 {
    g.playerY -= speed
}
if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.playerY < screenHeight - playerHeight {
    g.playerY += speed
}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(image.Black)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerX, g.playerY)
	screen.DrawImage(g.player, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func main() {
	game := &Game{
		playerX: 100,
		playerY: 100,
		player:  loadImage("player.png"),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Movimento do Jogador")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}