package compile

import (
	"DistroJudge/api"
	"context"
)

var languageMap map[api.Task_Language]Compile

type Result struct {
	OutPath string
	Time    uint64
	Memory  uint64
}

type Compile interface {
	// Compile is to compile the file into an executable file and return a compilation error
	Compile(code string, language api.Task_Language) (string, error)

	// Run is the file after executing Compile and returns the operation result.
	Run(ctx context.Context, path string, language api.Task_Language, in string, t, memory uint64) (*Result, error)
}

type Core struct {
}

func (c Core) Compile(code string, language api.Task_Language) (string, error) {
	return languageMap[language].Compile(code, language)
}

func (c Core) Run(ctx context.Context, path string, language api.Task_Language, in string, t, memory uint64) (*Result, error) {
	return languageMap[language].Run(ctx, path, language, in, t, memory)
}

func init() {
	languageMap = make(map[api.Task_Language]Compile)

	languageMap[api.Task_C] = &c_Compile{}
	languageMap[api.Task_JAVA] = &java_Compile{}
	languageMap[api.Task_PYTHON] = &python_Compile{}
	languageMap[api.Task_GOLANG] = &go_Compile{}
}
