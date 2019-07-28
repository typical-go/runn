package runn

// Executor do the code statement execution
type Executor struct {
	StopWhenError bool
}

// Execute all statement
func (e Executor) Execute(stmts ...interface{}) (err error) {
	var errs Errors
	for _, stmt := range stmts {

		switch stmt.(type) {
		case error:
			stmtErr := stmt.(error)
			if e.StopWhenError {
				return stmtErr
			}
			errs.Add(stmtErr)
		}
	}
	if len(errs) > 0 {
		err = errs
	}
	return
}
