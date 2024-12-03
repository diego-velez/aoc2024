package main

import (
	"os"
)

func ParseFile(f string, fileLineFunc func(string)) {
	file, err := os.Open(f)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	line := ""

	for {
		n, err := file.Read(buffer)
		if n > 0 {
			for _, b := range buffer[:n] {
				if b == '\n' {
					fileLineFunc(line)
					line = ""
				} else {
					line += string(b)
				}
			}
		}

		if err != nil {
			if line != "" {
				fileLineFunc(line)
			}
			break
		}
	}
}
