package result

import (
	"testing"
	"github.com/pkg/errors"
	"fmt"
)

func TestSSHResult_SSHResultToJson(t *testing.T) {
	s := &SSHResult{
		Err:errors.New("this is err"),
		Result:"result ...", ExitCode:-1,
	}

	ss, _ := s.ToJson()
	fmt.Println(ss)
	fmt.Println(string(ss))
}

func TestSFTPResult_ToJsonString(t *testing.T) {
	s := &SFTPResult{
		Err:errors.New("This is a err."),
		Changed:true,
	}

	ss := s.ToJsonString()
	fmt.Println(ss)
}