package interfaces

type Renderer interface {
	// TODO: simplify, Render(data interface{}) ([]byte,error)
	Render(data interface{}, templates []string, status int) []byte
	Status() int
	Headers() map[string]string
}
