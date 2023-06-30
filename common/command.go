package common

// Specify a new command
type Command struct {
	Usage func()
	Run   func([]string)
}
