package result

import (
	"testing"
	"github.com/pkg/errors"
	"fmt"

	"github.com/bypdhu/ssh-executor/common"
)

func TestSSHResult_SSHResultToJson(t *testing.T) {
	s := &SSHResult{
		Stdout:"result ...", ExitCode:-1,
	}

	ss, _ := s.ToJson()
	fmt.Println(ss)
	fmt.Println(string(ss))
}

func TestSFTPResult_ToJsonString(t *testing.T) {
	s := &SFTPResult{
		Changed:true,
	}

	ss := s.ToJsonString()
	fmt.Println(ss)
}

func TestBaseResult_ToJsonString(t *testing.T) {
	s := &BaseResult{
		SSHResult:SSHResult{Stdout:"this is result...", ExitCode:-1},
		SFTPResult:SFTPResult{Changed:true},
		Err:errors.New("This is a err."),
	}

	ss := s.ToJsonString(common.MODULE_COPY.String())
	fmt.Println(ss)

	ss = s.ToJsonString(common.MODULE_SHELL.String())
	fmt.Println(ss)
}