package keyMap

import "github.com/charmbracelet/bubbles/key"

const (
	SymbolicUpKey    = "↑"
	SymbolicDownKey  = "↓"
	SymbolicLeftKey  = "←"
	SymbolicRightKey = "→"
)

type KeyMapItem struct {
	keyBinding key.Binding
	// The key of the key item, not the keybinding. Example: "Up", "Quit"
	key string
	// Include in short help
	short bool
	// Include in long help
	long bool
	// Organize long help into multiple columns
	longCol int
	// Symbolic representation of key
	symbolicKey string
}

type KeyMap struct {
	Items []KeyMapItem
}

func NewKeyMap(items []KeyMapItem) KeyMap {
	return KeyMap{items}
}

func (k KeyMap) ShortHelp() []key.Binding {
	var shortItems []key.Binding
	for _, item := range k.Items {
		if item.short {
			shortItems = append(shortItems, item.keyBinding)
		}
	}
	return shortItems
}

// TEST: test this to make sure it can accept the longCol property in any order without blowing up.
func (k KeyMap) FullHelp() [][]key.Binding {
	var longItems [][]key.Binding
	for _, item := range k.Items {
		if item.long {
			if len(longItems) < item.longCol {
				longItems[item.longCol] = []key.Binding{item.keyBinding}
			} else {
				longItems[item.longCol] = append(longItems[item.longCol], item.keyBinding)
			}
		}
	}
	return longItems
}
