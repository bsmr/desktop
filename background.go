package desktop

import (
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"

	wmtheme "fyne.io/desktop/theme"
)

func wallpaperPath() string {
	pathEnv := Instance().Settings().Background()
	if pathEnv == "" {
		return ""
	}

	if stat, err := os.Stat(pathEnv); os.IsNotExist(err) || !stat.Mode().IsRegular() {
		return ""
	}

	return pathEnv
}

func newBackground() fyne.CanvasObject {
	imagePath := wallpaperPath()
	if imagePath != "" {
		return canvas.NewImageFromFile(imagePath)
	}
	return canvas.NewImageFromResource(wmtheme.Background)
}
