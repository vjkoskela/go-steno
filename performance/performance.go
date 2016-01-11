package main

import (
	"errors"
	"os"
	"github.com/pborman/uuid"
	"github.com/Sirupsen/logrus"
	"github.com/vjkoskela/gosteno"
	"time"
	"fmt"
)

// Execute this using the command: ./performance > /dev/null
//
// The result is a single line of output to standard error with the elapsed seconds.
func main() {
	// Configure logrus
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(new(gosteno.Formatter))
	logrus.SetLevel(logrus.DebugLevel)

	// Create steno logger
	var logger *gosteno.Logger = gosteno.GetLogger("performance.main")

	// Precreate output data
	var err error = errors.New("This is an error")
	var uuid1 string = uuid.New()
	var uuid2 string = uuid.New()

	// Execute test
	var iterations int = 100000
	var start time.Time = time.Now()
	for i := 0; i < iterations; i++ {
		logger.InfoBuilder().
		SetEvent("performance_event").
		SetMessage("This is a message from the steno logger").
		SetError(err).
		AddContext("requestId", uuid1).
		AddData("userId", uuid2).
		Log()
	}
	var end time.Time = time.Now()
	var elapsed float64 = float64(end.Sub(start).Nanoseconds())
	os.Stderr.WriteString(fmt.Sprintf("Elapsed %f seconds\n", elapsed / 1000000000.0))
}
