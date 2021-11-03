package interface_tese

import "testing"

type Programmer interface {
	WriteHelloWord() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWord() string {
	return "fmt.Println(\"Hello world\")"
}

func TestClient(t *testing.T) {
	var p Programmer = new(GoProgrammer)
	t.Log(p.WriteHelloWord())
}
