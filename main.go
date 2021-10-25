package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer stop()

	_, err := readLine()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	res := doSmth(ctx)
	fmt.Print(res)
}

func doSmth(ctx context.Context) string {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	select {
		case <- ctx.Done():
			return "ctx cancelled"
		case <- ticker.C:
			return "result"
	}
}

func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter something: ")
	line, err := reader.ReadString('\n')
	return strings.TrimSuffix(line, "\n"), err
}