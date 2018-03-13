package transform

type TransformConfig struct {
	Fim struct {
		AppGreylist []string
		FileList    []string
		SyscallList []string
		UserList    []int
		FieldLists  []struct {
			Pattern string
			Type    string
		}
	}
}
