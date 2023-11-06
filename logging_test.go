package logger

import (
	"bufio"
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestInitLogger(t *testing.T) {
	tests := []struct {
		name        string
		namespace   string
		development bool
		filepath    string
		level       zapcore.Level
	}{
		{
			name:        "Test without file",
			namespace:   "test",
			development: true,
			filepath:    "",
			level:       zapcore.InfoLevel,
		},

		{
			name:        "Test with file",
			namespace:   "test",
			development: true,
			filepath:    "log.txt",
			level:       zapcore.InfoLevel,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if err := InitLogger(Config{
				Namespace:   testCase.namespace,
				Development: testCase.development,
				Filepath:    testCase.filepath,
				Level:       testCase.level,
			}); err != nil {
				t.Error(err)
			}

			logger.Info("test")
			logger.Warn("test")

			if len(testCase.filepath) != 0 {
				file, err := os.Open(testCase.filepath)
				if err != nil {
					t.Error(err)
				}

				scanner := bufio.NewScanner(file)
				scanner.Split(bufio.ScanLines)

				for scanner.Scan() {
					fmt.Printf("line from log file: %s\n", scanner.Text())
				}

				_ = os.Remove(testCase.filepath)
			}
		})
	}
}
