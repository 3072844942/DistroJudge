package compile

import (
	"DistroJudge/api"
	"DistroJudge/sandbox"
	"context"
)

type c_Compile struct {
	//Compile
}

func (c c_Compile) Compile(code string, language api.Task_Language) (string, error) {
	return "", nil
}

func (c c_Compile) Run(ctx context.Context, path string, language api.Task_Language, in string, t, memory uint64) (*Result, error) {
	cmd := sandbox.Command(path)

	// todo
	if err := cmd.Run(); err != nil {
	}
	return nil, nil
}
