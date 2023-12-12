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

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("cmd> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
	
		line = strings.TrimSpace(line)
		line = strings.Trim(line, "\n")
		if strings.EqualFold("exit", line) {
			break
		}
		input := strings.Split(line, " ")
		executeInput(input)
	}
}

func executeInput(input []string) {
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