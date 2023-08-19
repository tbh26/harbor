package code

import (
	"fmt"
	"os"
)

func FilesDemo() {
	fmt.Println("Hello code.FilesDemo() world! (ch7)")

	FilesIntro()
	CreateFile()
	CreateAndSeekDemo()

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

func listDir(dirPath string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		panic(fmt.Sprintf("read directory %s failure; %s", dirPath, err))
	}

	fmt.Printf(" list %s \n", dirPath)
	for _, entry := range entries {
		filePath := fmt.Sprintf("%s/%s", dirPath, entry.Name())
		//fmt.Println(fullPath)
		fi, err := os.Stat(filePath)
		if err != nil {
			fmt.Printf("stat %s failure; %s \n", filePath, err)
		} else {
			fmt.Printf(" %s  %8d  %s \n", fi.Mode(), fi.Size(), fi.Name())
		}
	}
}

func deleteEntry(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(fmt.Sprintf("delete %s failure; %s", path, err))
	} else {
		fmt.Printf("%s deleted \n", path)
	}

}

func CreateFile() {
	fmt.Println("\n=-= create file demo =-=")
	tmpDir := "/var/tmp"

	filePath := fmt.Sprintf("%s/message-created", tmpDir)
	msg := []string{
		"Rule", "the", "world", "with", "Go!!!"}

	f, err := os.Create(filePath)
	if err != nil {
		panic(fmt.Sprintf("create %s failure; %s", filePath, err))
	}
	defer f.Close()

	fmt.Printf(" | %s | \n", filePath)
	for _, s := range msg {
		f.WriteString(s + "\n")
		fmt.Println(s)
	}

	listDir(tmpDir)
	deleteEntry(filePath)
	listDir(tmpDir)

	fmt.Println()
}

func CreateAndSeekDemo() {
	fmt.Println("\n=-= create and seek file demo =-=")

	//tmpDir := os.TempDir() // trailing /
	//fullPath := fmt.Sprintf("%smessage-create-and-seek", tmpDir)
	tmpDir := "/var/tmp"
	fullPath := fmt.Sprintf("%s/message-create-and-seek", tmpDir)

	file, err := os.Create(fullPath)
	if err != nil {
		panic(fmt.Sprintf("create %s failure; %s", fullPath, err))
	}
	defer func(f *os.File) {
		ce := f.Close()
		if ce != nil {
			fmt.Printf("failure on defered close %s; %s", f.Name(), ce)
		}
	}(file)
	//defer file.Close()

	msg := "Save the world with Go!!!"
	_, err = file.WriteString(msg)
	if err != nil {
		panic(fmt.Sprintf("write %s failure; %s", fullPath, err))
		//panic(err)
	}

	positions := []int{4, 10, 20}
	for _, i := range positions {
		_, err := file.Seek(int64(i), 0)
		if err != nil {
			panic(err)
		}
		file.Write([]byte("X"))
	}
	// Reset
	file.Seek(0, 0)
	// Read the result
	result := make([]byte, len(msg))
	_, err = file.Read(result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result)
	//
	listDir(tmpDir)
	deleteEntry(fullPath)
	listDir(tmpDir)

	fmt.Println()
}
