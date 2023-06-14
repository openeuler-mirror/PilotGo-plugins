package common

type MachineNode struct {
	UUID       string
	Department string
	IP         string
	CPUArch    string
	OSInfo     string
	State      int
}

type Batch struct {
	BatchUUID     string
	DepartmentIDs []string
	MachineUUIDs  []string
}
