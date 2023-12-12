package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	fmt.Print("cmd> ")

	reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }

	line = strings.TrimSpace(line)
	line = strings.Trim(line, "\n")

	input := strings.Split(line, " ")

	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		defer wg.Done()

		cmd := exec.Command(input[0])
		if len(input) > 1 {
			cmd = exec.Command(input[0], input[1:]...)
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		err = cmd.Wait()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}()
	wg.Wait()
}