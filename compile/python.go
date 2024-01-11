package compile

import (
	"DistroJudge/api"
	"context"
)

type python_Compile struct {
	//Compile
}

func (c python_Compile) Compile(code string, language api.Task_Language) (string, error) {
	return "", nil
}

func (c python_Compile) Run(ctx context.Context, path string, language api.Task_Language, in string, t, memory uint64) (*Result, error) {
	return nil, nil
}
