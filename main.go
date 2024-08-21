package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type status int

const (
    todo status = iota
    inProgress
    done
)

type Task struct {
    status status
    title string
    description string
}

func (t Task) FilterValue() string {
    return t.title
}

func (t Task) Title() string {
    return t.title
}

func (t Task) Description() string {
    return t.description
}

type Model struct {
    list list.Model
    err error
}

func New() *Model {
    return &Model{}
}

func (m *Model) initList(width, height int) {
    m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
    m.list.Title = "To Do"
    m.list.SetItems([]list.Item{
        Task{status: todo, title: "Item 1", description: "todo item number 1"},
        Task{status: todo, title: "Item 2", description: "todo item number 2"},
        Task{status: todo, title: "Item 3", description: "todo item number 3"},
    })
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.initList(msg.Width, msg.Height)
    }
    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m Model) View() string {
    return m.list.View()
}

func main() {
    m := New()
    p := tea.NewProgram(m)
    if _, err := p.Run(); err != nil {
        fmt.Println()
        os.Exit(1)
    }
}
