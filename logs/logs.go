package logs

import (
	"time"
	"github.com/Baligul/election/lib/file"
)

// WriteLogs function will be called in main package to automatically
// write errors to the log file at logFilePath
func WriteLogs(data string) error {

	const layout = "Jan 2, 2006 at 3:04pm (MST)"

	// Get log file
	logFile, err := file.Open("logs/logs.txt")
	defer file.Close(logFile)

	if err != nil {
		return err
	}

	data = time.Now().Format(layout) + " - " + data + "\n"
	_, err = file.Write(logFile, []byte(data))

	if err != nil {
		return err
	}

    err = file.Sync(logFile)
	if err != nil {
		return err
	}
    
	return nil
}