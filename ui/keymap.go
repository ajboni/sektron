package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
)

// TODO: implement left and right for parameters
var (
	stepSelectKeys  = []string{"a", "z", "e", "r", "t", "y", "u", "i", "q", "s", "d", "f", "g", "h", "j", "k"}
	stepToggleKeys  = []string{"A", "Z", "E", "R", "T", "Y", "U", "I", "Q", "S", "D", "F", "G", "H", "J", "K"}
	trackSelectKeys = []string{"&", "é", "\"", "'", "(", "-", "è", "_", "ç", "à"}
	trackToggleKeys = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
)

type keyMap struct {
	TogglePlay key.Binding
	Mode       key.Binding

	Add    key.Binding
	Remove key.Binding

	StepSelectIndex map[string]int
	StepSelect      key.Binding

	StepToggleIndex map[string]int
	StepToggle      key.Binding

	TrackSelectIndex map[string]int
	TrackSelect      key.Binding

	TrackToggleIndex map[string]int
	TrackToggle      key.Binding

	TrackPageUp   key.Binding
	TrackPageDown key.Binding

	TempoUp       key.Binding
	TempoDown     key.Binding
	TempoFineUp   key.Binding
	TempoFineDown key.Binding

	ParamSelectLeft  key.Binding
	ParamSelectRight key.Binding

	ParamValueUp   key.Binding
	ParamValueDown key.Binding

	Help key.Binding
	Quit key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.TogglePlay, k.Mode, k.Add, k.Remove, k.TempoUp, k.TempoDown, k.TempoFineUp, k.TempoFineDown},
		{k.StepSelect, k.StepToggle, k.TrackSelect, k.TrackToggle, k.TrackPageUp, k.TrackPageDown, k.ParamValueUp, k.ParamValueDown},
		{k.Help, k.Quit},
	}
}

// DefaultKeyMap returns the default key mapping.
func DefaultKeyMap() keyMap {
	km := keyMap{
		TogglePlay: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "toggle play"),
		),
		Mode: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "toggle mode (track, record)"),
		),
		Add: key.NewBinding(
			key.WithKeys("="),
			key.WithHelp("=", "add track|step"),
		),
		Remove: key.NewBinding(
			key.WithKeys(")"),
			key.WithHelp(")", "remove track|step"),
		),
		StepSelectIndex: map[string]int{},
		StepSelect: key.NewBinding(
			key.WithKeys(stepSelectKeys...),
			key.WithHelp(strings.Join(stepSelectKeys, "/"), "select track|step 1 to 16"),
		),
		StepToggleIndex: map[string]int{},
		StepToggle: key.NewBinding(
			key.WithKeys(stepToggleKeys...),
			key.WithHelp(strings.Join(stepToggleKeys, "/"), "toggle track|step 1 to 16"),
		),
		TrackSelectIndex: map[string]int{},
		TrackSelect: key.NewBinding(
			key.WithKeys(trackSelectKeys...),
			key.WithHelp(strings.Join(trackSelectKeys, "/"), "select track 1 to 10"),
		),
		TrackToggleIndex: map[string]int{},
		TrackToggle: key.NewBinding(
			key.WithKeys(trackToggleKeys...),
			key.WithHelp(strings.Join(trackToggleKeys, "/"), "select track 1 to 10"),
		),
		TrackPageUp: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "track page up"),
		),
		TrackPageDown: key.NewBinding(
			key.WithKeys("m"),
			key.WithHelp("m", "track page down"),
		),
		TempoUp: key.NewBinding(
			key.WithKeys("pgup"),
			key.WithHelp("page up", "tempo up (1 bpm)"),
		),
		TempoDown: key.NewBinding(
			key.WithKeys("pgdown"),
			key.WithHelp("page down", "tempo down (1 bpm)"),
		),
		TempoFineUp: key.NewBinding(
			key.WithKeys("alt+pgup"),
			key.WithHelp("alt+page up", "tempo up (0.1 bpm)"),
		),
		TempoFineDown: key.NewBinding(
			key.WithKeys("alt+pgdown"),
			key.WithHelp("alt+page down", "tempo down (0.1 bpm)"),
		),
		ParamSelectLeft: key.NewBinding(
			key.WithKeys("left"),
			key.WithHelp("←", "parameter select left"),
		),
		ParamSelectRight: key.NewBinding(
			key.WithKeys("right"),
			key.WithHelp("→", "parameter select left"),
		),
		ParamValueUp: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "increase selected parameter value"),
		),
		ParamValueDown: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "decrease selected parameter value"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "esc"),
			key.WithHelp("ctrl+c/esc", "quit"),
		),
	}
	for i, k := range stepSelectKeys {
		km.StepSelectIndex[k] = i
	}
	for i, k := range stepToggleKeys {
		km.StepToggleIndex[k] = i
	}
	for i, k := range trackSelectKeys {
		km.TrackSelectIndex[k] = i
	}
	for i, k := range trackToggleKeys {
		km.TrackToggleIndex[k] = i
	}
	return km
}
