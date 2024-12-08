package general_select_with_desc

import "github.com/charmbracelet/bubbles/list"

func WithStatusHidden(l *list.Model) {
	l.SetShowStatusBar(false)
}
