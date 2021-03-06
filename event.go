package tui

import termbox "github.com/nsf/termbox-go"

type Key termbox.Key
type EventType termbox.EventType

const (
	EventKey EventType = iota
	EventResize
	EventMouse
	EventError
	EventInterrupt
	EventRaw
	EventNone
)

type Event struct {
	Type EventType
	Key  Key
	Ch   rune
}

const (
	KeyEnter      = Key(termbox.KeyEnter)
	KeySpace      = Key(termbox.KeySpace)
	KeyTab        = Key(termbox.KeyTab)
	KeyEsc        = Key(termbox.KeyEsc)
	KeyBackspace  = Key(termbox.KeyBackspace)
	KeyBackspace2 = Key(termbox.KeyBackspace2)
	KeyArrowUp    = Key(termbox.KeyArrowUp)
	KeyArrowDown  = Key(termbox.KeyArrowDown)
	KeyArrowLeft  = Key(termbox.KeyArrowLeft)
	KeyArrowRight = Key(termbox.KeyArrowRight)
)

func convertEvent(tev termbox.Event) Event {
	return Event{
		Type: EventType(tev.Type),
		Key:  Key(tev.Key),
		Ch:   tev.Ch,
	}
}
