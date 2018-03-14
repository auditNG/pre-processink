package transform

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	err := f.Init(config)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Instantiate transform helper
	transformHelper := NewTransformHelper()

	//Get syscall
	syscall, err := transformHelper.GetIntValue(message, "syscall=")
	if nil != err {
		fmt.Println("Unable to get syscall")
		return nil
	}
	fmt.Println(syscall)

	//Get exitcode
	exitcode, err := transformHelper.GetIntValue(message, "exit=")
	if nil != err {
		fmt.Println("Unable to get exitcode")
		return nil
	}
	fmt.Println(exitcode)

	//Get executable
	executable, err := transformHelper.GetStringValue(message, "exe=")
	if nil != err {
		fmt.Println("Unable to get executable")
		return nil
	}
	fmt.Println(executable)

	//Get user
	user, err := transformHelper.GetIntValue(message, "uid=")
	if nil != err {
		fmt.Println("Unable to get user id")
		return nil
	}
	fmt.Println("user: " + string(user))

	//Label this event into RED/YELLOW/GREEN
	label := f.applyLabelAlgo(message, syscall, exitcode, executable, user)

	outputLine := f.constructOutputLine(syscall, exitcode, executable, user, label)

	if _, err := outputFile.Write([]byte(outputLine)); err != nil {
		fmt.Println("Error weiting line to output file")
		fmt.Println(err)
	}
	return nil
}

func (f FIMTransformer) constructOutputLine(syscall int, exitcode int, executable string, user int, label int) string {
	logLine := []string{strconv.Itoa(syscall), strconv.Itoa(exitcode), executable, strconv.Itoa(user), strconv.Itoa(label)}
	return strings.Join(logLine, ",") + "\n"
}

func (f FIMTransformer) isUserInWatchList(user int) bool {
	for _, val := range f.confObj.Fim.UserList {
		if user == val {
			return true
		}
	}
	return false
}

func (f FIMTransformer) isSyscallInWatchList(syscall int) bool {
	for _, val := range f.confObj.Fim.SyscallList {
		if syscall == val {
			return true
		}
	}
	return false
}

func (f FIMTransformer) isFileInWatchList(message string) bool {
	//Get filename
	transformHelper := NewTransformHelper()
	for _, val := range f.confObj.Fim.FileList {
		filename, err := transformHelper.GetStringValue(message, val)
		if nil != err && "" != filename {
			fmt.Println("File found")
			return true
		}
	}
	return false
}

func (f FIMTransformer) isExeInWatchList(exe string) bool {
	for _, val := range f.confObj.Fim.AppGreylist {
		if strings.Contains(exe, val) {
			fmt.Println(val)
			return true
		}
	}
	return false
}

func (f FIMTransformer) applyLabelAlgo(message string, syscall int, exitcode int, executable string, user int) int {

	if f.isFileInWatchList(message) && f.isSyscallInWatchList(syscall) && f.isExeInWatchList(executable) && f.isUserInWatchList(user) {
		return RED
	}

	if f.isFileInWatchList(message) && f.isSyscallInWatchList(syscall) {
		return YELLOW
	}

	return GREEN
}
