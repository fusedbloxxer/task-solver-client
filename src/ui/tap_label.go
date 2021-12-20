package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TapLabel struct {
	widget.Label
	OnTap func()
}

func NewTappableLabel(message string, onTap func()) *TapLabel {
	label := &TapLabel{}
	label.ExtendBaseWidget(label)
	label.Text = message
	return label
}

func (tapLabel *TapLabel) Tapped(*fyne.PointEvent) {
	tapLabel.OnTap()
}

func (tapLabel *TapLabel) TappedSecondary(*fyne.PointEvent) {

}
