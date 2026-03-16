package todo

type Task struct {
	Name string
	Description string
	CreatedAt time.Time
	Completed bool
	CompletedAt time.Time
}

func NewTask(name, desription string) Task {
	return Task{
		Name: name,
		Description: desription,
		CreatedAt: time.Now(),
		Completed: false,
	}
}

