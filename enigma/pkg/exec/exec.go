package exec

import (
	"fmt"
	"os"

	"github.com/go-cmd/cmd"
)

func ExecCMD(command string, args ...string) {
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	envCmd := cmd.NewCmdOptions(cmdOptions, command, args...)

	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)

		for envCmd.Stdout != nil || envCmd.Stderr != nil {
			select {
			case line, open := <-envCmd.Stdout:
				if !open {
					envCmd.Stdout = nil
					continue
				}
				fmt.Println(line)
			case line, open := <-envCmd.Stderr:
				if !open {
					envCmd.Stderr = nil
					continue
				}
				fmt.Fprintln(os.Stderr, line)
			}
		}
	}()

	<-envCmd.Start()

	<-doneChan
}
