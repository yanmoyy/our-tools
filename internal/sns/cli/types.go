package cli

type Command[T any] struct {
	Name        string
	Description string
	Callback    func(cfg *T, args ...string) error
	Helper      func()
}

type Commands[T any] map[string]Command[T]
