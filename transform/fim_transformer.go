package transform

import (
	"encoding/json"
	"fmt"
	"os"
	// "strconv"
	// "strings"
)

const (
	GREEN = 1 + iota
	YELLOW
	RED
)

func NewFIMTransformer() FIMTransformer {
	return FIMTransformer{
		confObj: new(TransformConfig),
	}
}

type FIMTransformer struct {
	confObj *TransformConfig
}

func (f FIMTransformer) Init(config string) error {
	err := json.Unmarshal([]byte(config), f.confObj)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (f FIMTransformer) Process(message string, config string, outputFile *os.File) error {
	// Read in config
	// fmt.Println(message[0])
	err := f.Init(config)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

	// Instantiate transform helper
	//transformHelper := NewTransformHelper()

	// name,_:=message["name"]
	// version,_:=message["version"]
	// major,_:=message["major"]
	// minor,_:=message["minor"]


	//Label this event into RED/YELLOW/GREEN
	// f.applyLabel(name,version)
	// f.applyLabel(name,version,major)
	// label := f.applyLabelAlgo(name, name, version, major, minor)

	// outputLine := f.constructOutputLine(name, version, major, minor, label)

// 	if _, err := outputFile.Write([]byte(outputLine)); err != nil {
// 		fmt.Println("Error weiting line to output file")
// 		fmt.Println(err)
// 	}
// 	return nil
// }
//
// func (f FIMTransformer) constructOutputLine(syscall string, exitcode string, executable string, user string, label int) string {
// 	logLine := []string{syscall, exitcode, executable, user, strconv.Itoa(label)}
// 	return strings.Join(logLine, ",") + "\n"
// }
//
// func (f FIMTransformer) isUserInWatchList(user string) bool {
// 	for _,val := range f.confObj.Fim.UserList {
// 		fmt.Println(val)
// 		// if user == val {
// 		// 	return true
// 		// }
// 	}
// 	return false
// }
//
// func (f FIMTransformer) isSyscallInWatchList(syscall string) bool {
// 	for _, val := range f.confObj.Fim.SyscallList {
// 		if syscall == val {
// 			return true
// 		}
// 	}
// 	return false
// }
//
//
//
// func (f FIMTransformer) isFileInWatchList(message string) bool {
// 	//Get filename
// 	transformHelper := NewTransformHelper()
// 	for _, val := range f.confObj.Fim.FileList{
//
// 		filename, err := transformHelper.GetStringValue(message, val)
// 		if nil == err && "" != filename {
// 			fmt.Println("File found")
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func (f FIMTransformer) isExeInWatchList(exe string) bool {
// 	for _, val := range f.confObj.Fim.AppGreylist {
// 		if strings.Contains(exe, val) {
// 			fmt.Println(val)
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func (f FIMTransformer) applyLabelAlgo(message string, syscall string, exitcode string, executable string, user string) int {
//
// 	if f.isFileInWatchList(message) && f.isSyscallInWatchList(syscall) && f.isExeInWatchList(executable) && f.isUserInWatchList(user) {
// 		fmt.Println("RED")
// 		return RED
// 	}
//
// 	if f.isFileInWatchList(message) && f.isSyscallInWatchList(syscall) {
// 		fmt.Println("Yellow")
// 		return YELLOW
// 	}
//
// 	return GREEN
// }
//
//
// func (f FIMTransformer) applyLabel(args ...interface{}) {
//
// }
