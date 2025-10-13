package model

type CmdResult struct {
	RetCode int
	Stdout  string
	Stderr  string
}

type ScriptsRun struct {
	ScriptType    string
	ScriptContent string
	Params        string
	TimeOut       int
}
