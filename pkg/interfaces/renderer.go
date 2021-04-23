package interfaces

type Renderer interface {
	Render(data interface{}, templates []string, status int) []byte
	Status() int
	Headers() map[string]string
}
