package nerve

import (
	"github.com/n0rad/go-erlog/errs"
	"os/exec"
)

type CheckExec struct {
	CheckCommon
	Command []string
}

func NewCheckExec() *CheckExec {
	return &CheckExec{}
}

func (x *CheckExec) Init(s *Service) error {
	if len(x.Command) == 0 {
		return errs.With("Exec command type require a command")
	}
	x.fields = x.fields.WithField("command", x.Command)
	return nil
}

func (x *CheckExec) Check() error {
	command := exec.Command(x.Command[0], x.Command[1:]...)
	data, err := command.CombinedOutput()
	if err != nil {
		return errs.WithEF(err, x.fields.WithField("output", data), "Command failed")
	}
	return nil
}