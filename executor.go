package runn

import (
	"bytes"
	"fmt"
	"os/exec"
)

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
		case Runner:
			runner := stmt.(Runner)
			runErr := runner.Run()
			if runErr != nil {
				if e.StopWhenError {
					return runErr
				}
				errs.Add(runErr)
			}
		case *exec.Cmd:
			execCmd := stmt.(*exec.Cmd)
			buf := bytes.Buffer{}
			execCmd.Stdout = &buf
			execCmd.Stderr = &buf

			runErr := execCmd.Run()
			if err != nil {
				runErr = fmt.Errorf("%s: %s", err.Error(), buf.String())
				if e.StopWhenError {
					return runErr
				}
				errs.Add(runErr)
			}
		}
	}
	if len(errs) > 0 {
		err = errs
	}
	return
}
