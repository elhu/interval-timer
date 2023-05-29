package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	currentTimer int
	timers       []time.Duration
	elapsed      time.Duration
}

type TickMsg time.Time

func initialModel() model {
	return model{
		currentTimer: 0,
		timers: []time.Duration{
			3 * time.Second,
			5 * time.Second,
			7 * time.Second,
		},
	}
}

func doTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return doTick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		}
	case TickMsg:
		m.elapsed += time.Second
		if m.elapsed >= m.timers[m.currentTimer] {
			m.currentTimer++
			m.elapsed = 0
			if m.currentTimer == len(m.timers) {
				m.currentTimer = 0
			}
		}
		return m, doTick()
	}
	return m, nil
}

func (m model) View() string {
	s := "Timers:\n\n"
	for i, t := range m.timers {
		if m.currentTimer == i {
			s += fmt.Sprintf("> %s\n", (t - m.elapsed).String())
		} else {
			s += fmt.Sprintf("  %s\n", t.String())
		}
	}
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
