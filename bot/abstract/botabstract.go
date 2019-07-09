package abstract

type Handler interface {
	CanHandle(msg string) bool
	Handle() (string, error)
}

func FindCommand(commands []string, msg string) bool {
	for _,v := range commands {
		if v == msg{
			return true
		}
	}
	return false
}