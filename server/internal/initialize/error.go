package initialize

type InitDBError struct {
	msg string
}

func (e *InitDBError) Error() string {
	return e.msg
}
