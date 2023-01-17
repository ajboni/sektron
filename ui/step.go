package ui

import (
	"fmt"
	"sektron/sequencer"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

const (
	stepWidth  = 15
	stepHeight = stepWidth / 2
)

var (
	stepStyle = lipgloss.NewStyle().
			Margin(1, 2, 0, 0)
	stepActiveStyle = lipgloss.NewStyle().
			Margin(1, 0, 0, 0)
	stepVelocityStyle = lipgloss.NewStyle().
				Margin(1, 1, 0, 0)
	textStyle = lipgloss.NewStyle().
			Foreground(secondaryTextColor).
			Padding(1, 1, 1, 2).
			Bold(true)
)

func (m mainModel) renderStep(step sequencer.Step) string {
	content := m.renderStepContent(step)
	var stepStr string
	velocityIndicator := []string{}
	width, height := m.stepSize()

	var stepCurrentColor, stepActiveColor, stepInactiveColor lipgloss.Color
	if step.Track().IsActive() {
		stepCurrentColor = currentColor
		stepActiveColor = activeColor
		stepInactiveColor = inactiveColor
	} else {
		stepCurrentColor = currentDimmedColor
		stepActiveColor = activeDimmedColor
		stepInactiveColor = inactiveDimmedColor
	}

	if m.seq.IsPlaying() && step.IsCurrentStep() {
		stepStr = stepStyle.Render(lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			textStyle.Background(stepCurrentColor).Render(content),
			lipgloss.WithWhitespaceBackground(stepCurrentColor),
		))
	} else if step.IsActive() {
		stepStr = stepActiveStyle.Render(lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			textStyle.Background(stepActiveColor).Render(content),
			lipgloss.WithWhitespaceBackground(stepActiveColor),
		))

		velocityValue := int(127-step.Velocity()) * lipgloss.Height(stepStr) / 127
		for i := 1; i < lipgloss.Height(stepStr); i++ {
			if velocityValue < i {
				velocityIndicator = append(velocityIndicator, "█")
			} else {
				velocityIndicator = append(velocityIndicator, " ")
			}
		}
	} else {
		stepStr = stepStyle.Render(lipgloss.Place(
			width,
			height,
			lipgloss.Left,
			lipgloss.Top,
			textStyle.Background(stepInactiveColor).Render(content),
			lipgloss.WithWhitespaceBackground(stepInactiveColor),
		))
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		stepStr,
		stepVelocityStyle.Render(lipgloss.JoinVertical(lipgloss.Left, velocityIndicator...)),
	)
}

func (m mainModel) stepSize() (int, int) {
	width := m.width/stepsPerLine - 2
	height := width/2 - 1
	if width < stepWidth || height < stepHeight {
		return stepWidth, stepHeight
	}
	return width, height
}

func (m mainModel) renderStepContent(step sequencer.Step) string {
	activeText := ""
	if step.Position() == m.activeStep {
		activeText = "♦"
	}
	if !step.IsActive() {
		return strconv.Itoa(step.Position() + 1)
	}
	return lipgloss.JoinVertical(
		lipgloss.Left,
		fmt.Sprintf("%d%s", step.Position()+1, activeText),
		note(step.Chord()[0]).Display(),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			lipgloss.NewStyle().
				Render(fmt.Sprintf("%.1f/%d", float64(step.Length())/6.0, m.trackPagesNb()*stepsPerPage)),
			lipgloss.NewStyle().
				MarginLeft(2).
				Render(fmt.Sprintf("%d%%", step.Probability())),
		),
	)
}
