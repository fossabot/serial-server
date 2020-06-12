package server

type Logger interface {
	Println(v ...interface{})
}

type noopLogger struct{}

func (l noopLogger) Println(v ...interface{}) {}
