package model

type CmdResult struct {
	RetCode int
	Stdout  string
	Stderr  string
}

type ScriptsRun struct {
	JobId         string
	ScriptType    string
	ScriptContent string
	Params        string
	TimeOut       int
}
