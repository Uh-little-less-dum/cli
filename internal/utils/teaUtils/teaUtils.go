package teaUtils

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// func NoExitFilter() {
// 	noExit := func(m tea.Model) {

// 	}

func NoExitFilter(m tea.Model, msg tea.Msg) tea.Msg {
	if _, ok := msg.(tea.QuitMsg); !ok {
		return msg
	}

	return nil
}

type Accessor[T any] interface {
	Get() T
	Set(value T)
}

type EmbeddedAccessor[T any] struct {
	value T
}

// Get gets the value.
func (a *EmbeddedAccessor[T]) Get() T {
	return a.value
}

// Set sets the value.
func (a *EmbeddedAccessor[T]) Set(value T) {
	a.value = value
}

// PointerAccessor allows field value to be exposed as a pointed variable.
type PointerAccessor[T any] struct {
	value *T
}

// NewPointerAccessor returns a new pointer accessor.
func NewPointerAccessor[T any](value *T) *PointerAccessor[T] {
	return &PointerAccessor[T]{
		value: value,
	}
}

// Get gets the value.
func (a *PointerAccessor[T]) Get() T {
	return *a.value
}

// Set sets the value.
func (a *PointerAccessor[T]) Set(value T) {
	*a.value = value
}

type Eval[T any] struct {
	val T
	fn  func() T

	bindings     any
	bindingsHash uint64
	cache        map[uint64]T

	loading      bool
	loadingStart time.Time
}

// type Theme struct {
// 	Form           lipgloss.Style
// 	Group          lipgloss.Style
// 	FieldSeparator lipgloss.Style
// 	Blurred        FieldStyles
// 	Focused        FieldStyles
// 	Help           help.Styles
// }
