package ports

type Executors interface {
	ExecuteCommand(commands ...string)
}
