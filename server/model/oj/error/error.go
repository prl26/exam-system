package exception

type CompileError struct {
	Msg string
}

func (c CompileError) Error() string {
	//panic("implement me")
	return c.Msg
}
