package code

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func StdIoDemo() {
	fmt.Println("Hello code.StdIoDemo() world! (ch7)")

	StdIoIntroDemo()
	StdInDemo()
	DialogDemo()
	ScannerDemo()
	LastDemo()

	fmt.Println()
}

func StdIoIntroDemo() {
	fmt.Println("\n=-= standard IO intro/demo, standard out  =-=")

	msg := []byte("Save the world with Go.\n")
	n, err := os.Stdout.Write(msg)
	if err != nil {
		panic(fmt.Sprintf("write standard out failure; %s \n", err))
	}
	fmt.Printf(" - written %d characters\n", n)

	fmt.Println()
}

func StdInDemo() {
	fmt.Println("\n=-= standard in demo =-=")

	fmt.Printf("Please type some stuff: ")
	target := make([]byte, 50)
	n, err := os.Stdin.Read(target)
	if err != nil {
		panic(fmt.Sprintf("read standard in failure; %s \n", err))
	}
	msg := string(target[:n])
	fmt.Println(n, strings.ToUpper(msg))

	fmt.Println()
}

func DialogDemo() {
	fmt.Println("\n=-= dialog demo =-=")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">>> What do you have to say?\n")
	fmt.Print("<<< ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(">>> You're right!!!")
	fmt.Println(strings.ToUpper(text))

	fmt.Println()
}

func ScannerDemo() {
	fmt.Println("\n=-= scanner demo =-=")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(">>> What do you have to say?\n")
	counter := 0
	for scanner.Scan() {
		text := scanner.Text()
		counter = counter + len(text)
		if counter > 15 {
			fmt.Println("break!")
			break
		}
	}
	fmt.Println("that's enough")

	fmt.Println()
}

func LastDemo() {
	fmt.Println("\n=-= last demo =-=")

	writer := bufio.NewWriter(os.Stdout)

	msg := "Save the world with Go!!!"
	for _, letter := range msg {
		time.Sleep(time.Millisecond * 300)
		writer.WriteByte(byte(letter))
		writer.Flush()
	}

	fmt.Println()
}
