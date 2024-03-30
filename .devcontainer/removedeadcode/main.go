package main

import (
	"io/fs"
	"os"
	"os/exec"
	"strings"

	"main/detect_deadcode"
)

func main() {
	path := os.Args[1]
	exec.Command("gofmt", "-w", path).Run()
	ranges := detect_deadcode.DetectDeadCodeRanges(path)

	// read code
	codebyte, _ := os.ReadFile(path)
	code := string(codebyte)
	// fmt.Println(code)
	codelines := strings.Split(code, "\n")

	// Include function/method comments in ranges
	for i := 0; i < len(ranges); i++ {
		for ranges[i][0] > 0 && len(codelines[ranges[i][0]-1]) != 0 {
			ranges[i][0]--
		}
	}

	// construct final code
	j := 0
	result := []byte{}
	for i := 0; i < len(codelines); i++ {
		if j < len(ranges) && i == ranges[j][0] {
			i = ranges[j][1] + 1
			j++
		}
		if i < len(codelines) {
			result = append(result, []byte(codelines[i]+"\n")...)
		}
	}
	os.WriteFile(path, []byte(result), fs.ModeType)

	// run goimports on the final code
	exec.Command("goimports", "-w", path).Run()
}
