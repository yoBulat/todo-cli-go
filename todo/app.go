package todo

type App struct {
	Tasks map[string]Task
	Events []Event
}