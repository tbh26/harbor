package code

import (
	"fmt"
	"os"
)

func FilesDemo() {
	fmt.Println("Hello code.FilesDemo() world! (ch7)")

	FilesIntro()

	fmt.Println()
}

func FilesIntro() {
	fmt.Println("\n=-= FilesDemo =-=")

	msg := "Build systems with Go.\nSave the world.\n"
	//msg = "Build systems with Go.\nSave the world?" // no newline at end
	filePath := "/var/tmp/message"

	//err := ioutil.WriteFile(filePath, []byte(msg), 0644) // deprecated
	err := os.WriteFile(filePath, []byte(msg), 0644)
	if err != nil {
		panic(fmt.Sprintf("write %s failure; %s", filePath, err))
	}

	//filePath = "/a/non/existing/path/message"

	//read, err := ioutil.ReadFile(filePath) // deprecated
	read, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("read %s failure; %s", filePath, err))
	}
	newLine := byte('\n')
	lastByte := read[len(read)-1]
	if lastByte == newLine {
		fmt.Printf(" | %s |\n%s", filePath, read)
	} else {
		fmt.Printf(" | %s |\n%s\n", filePath, read)
	}

	err = os.Remove(filePath)
	if err != nil {
		panic(fmt.Sprintf("delete %s failure; %s", filePath, err))
	}

}
