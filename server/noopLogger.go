package server

type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type noopLogger struct{}

func (l noopLogger) Println(v ...interface{})               {}
func (l noopLogger) Printf(format string, v ...interface{}) {}
