package popup

import (
	"image/color"
	"time"
)

type Message struct {
	Expiry     time.Time
	Foreground color.Color
	Background color.Color
	Text       string
}
