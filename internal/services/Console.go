package services

import (
	"bufio"
	"fmt"
	"os/exec"
)

func Call(fileName, songUrl, format string) error {
	command := fmt.Sprintf(
		"yt-dlp --cookies cookies.txt -f \"bestaudio\" "+
			"-P D:\\JavaContent\\IdeaProjects\\MusicDealerWin\\resources\\Audio\\downloads\\ -o %s.%s %s",
		fileName, format, songUrl,
	)
	fmt.Println("Executing command:", command)

	cmd := exec.Command("cmd", "/C", command)
	cmd.Dir = `D:\JavaContent\IdeaProjects\MusicDealerWin\resources\Audio`

	cmd.Stderr = cmd.Stdout

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	reader := bufio.NewReader(stdout)
	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err.Error() != "EOF" {
					fmt.Printf("Error reading stdout: %v\n", err)
				}
				break
			}
			fmt.Print(line)
		}
	}()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command finished with error: %w", err)
	}

	fmt.Println("Command executed successfully")
	return nil
}
