/*
Copyright 2016 Ville Koskela

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gosteno

import (
	"io"
	"github.com/Sirupsen/logrus"
)

var (
	noopLogBuilder LogBuilder = new(NoOpLogBuilder)
)

// Logger implementation. Interface compatible with standard Go log and github.com/Sirupsen/logrus. However, the power
// of the implementation is in the methods returning LogBuilder instances. The "Print" methods are mapped to the info
// level. As with other implementations fatal will exit and panic will halt.
type Logger struct {
	name string
	logger *logrus.Logger
}

func NewLogger(n string, l *logrus.Logger) *Logger {
	if (l == nil) {
		l = logrus.StandardLogger()
	}
	return &Logger{name: n, logger: l}
}

// ** Log Builder **

// Debug with LogBuilder. Recommended.
func (l *Logger) DebugBuilder() LogBuilder {
	if l.logger.Level >= logrus.DebugLevel {
		return createLogBuilder(l.logger, logrus.DebugLevel, l.name)
	} else {
		return noopLogBuilder
	}
}

// Info with LogBuilder. Recommended.
func (l *Logger) InfoBuilder() LogBuilder {
	if l.logger.Level >= logrus.InfoLevel {
		return createLogBuilder(l.logger, logrus.InfoLevel, l.name)
	} else {
		return noopLogBuilder
	}
}

// Warn with LogBuilder. Recommended.
func (l *Logger) WarnBuilder() LogBuilder {
	if l.logger.Level >= logrus.WarnLevel {
		return createLogBuilder(l.logger, logrus.WarnLevel, l.name)
	} else {
		return noopLogBuilder
	}
}

// Warning with LogBuilder. Recommended.
func (l *Logger) WarningBuilder() LogBuilder {
	return l.WarnBuilder()
}

// Error with LogBuilder. Recommended.
func (l *Logger) ErrorBuilder() LogBuilder {
	if l.logger.Level >= logrus.ErrorLevel {
		return createLogBuilder(l.logger, logrus.ErrorLevel, l.name)
	} else {
		return noopLogBuilder
	}
}

// Fatal with LogBuilder. Recommended. This implementation like the standard library causes the program to exit.
func (l *Logger) FatalBuilder() LogBuilder {
	if l.logger.Level >= logrus.FatalLevel {
		return createLogBuilder(l.logger, logrus.FatalLevel, l.name)
	} else {
		return noopLogBuilder
	}
}

// Panic with LogBuilder. Recommended. This implementation like the standard library causes the program to panic.
func (l *Logger) PanicBuilder() LogBuilder {
	if l.logger.Level >= logrus.PanicLevel {
		return createLogBuilder(l.logger, logrus.PanicLevel, l.name)
	} else {
		return noopLogBuilder
	}
}

// ** Go Log Compatibility **

// Print from standard Go log library. Provided for compatibility.
func (l *Logger) Print(args ...interface{}) {
	if l.logger.Level >= logrus.InfoLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Info()
	}
}

// Printf from standard Go log library. Provided for compatibility.
func (l *Logger) Printf(format string, args ...interface{}) {
	if l.logger.Level >= logrus.InfoLevel {
		l.logger.Infof(format, args...)
	}
}

// Println from standard Go log library. Provided for compatibility.
func (l *Logger) Println(args ...interface{}) {
	if l.logger.Level >= logrus.InfoLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Info()
	}
}

// Panic from standard Go log library. This implementation like the standard library causes the program to panic.
// Provided for compatibility.
func (l *Logger) Panic(args ...interface{}) {
	if l.logger.Level >= logrus.PanicLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Panic()
	}
}

// Panicf from standard Go log library. This implementation like the standard library causes the program to panic.
// Provided for compatibility.
func (l *Logger) Panicf(format string, args ...interface{}) {
	if l.logger.Level >= logrus.PanicLevel {
		l.logger.Panicf(format, args...)
	}
}

// Panicln from standard Go log library. This implementation like the standard library causes the program to panic.
// Provided for compatibility.
func (l *Logger) Panicln(args ...interface{}) {
	if l.logger.Level >= logrus.PanicLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Panic()
	}
}

// Fatal from standard Go log library. This implementation like the standard library causes the program to exit.
// Provided for compatibility.
func (l *Logger) Fatal(args ...interface{}) {
	if l.logger.Level >= logrus.FatalLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Fatal()
	}
}

// Fatalf from standard Go log library. This implementation like the standard library causes the program to exit.
// Provided for compatibility.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	if l.logger.Level >= logrus.FatalLevel {
		l.logger.Fatalf(format, args...)
	}
}

// Fatalln from standard Go log library. This implementation like the standard library causes the program to exit.
// Provided for compatibility.
func (l *Logger) Fatalln(args ...interface{}) {
	if l.logger.Level >= logrus.FatalLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Fatal()
	}
}

// Output from standard Go log library. This is mapped to Info without the stack trace. Provided for compatibility.
func (l *Logger) Output(calldepth int, s string) error {
	l.Info(s)
	return nil
}

// Flags from standard Go log library. This is a no-op. Provided for compatibility.
func (l *Logger) Flags() int {
	return 0;
}

// SetFlags from standard Go log library. This is a no-op. Provided for compatibility.
func (l *Logger) SetFlags(flag int) {
	// No-op
}

// Prefix from standard Go log library. This is a no-op. Provided for compatibility.
func (l *Logger) Prefix() string {
	return ""
}

// SetPrefix from standard Go log library. This is a no-op. Provided for compatibility.
func (l *Logger) SetPrefix(prefix string) {
	// No-op
}

// SetOutput from standard Go log library. This is a no-op. Provided for compatibility.
func SetOutput(w io.Writer) {
	// No-op
}

// Flags from standard Go log library. This is a no-op. Provided for compatibility.
func Flags() int {
	return 0;
}

// SetFlags from standard Go log library. This is a no-op. Provided for compatibility.
func SetFlags(flag int) {
	// No-op
}

// Prefix from standard Go log library. This is a no-op. Provided for compatibility.
func Prefix() string {
	return ""
}

// SetPrefix from standard Go log library. This is a no-op. Provided for compatibility.
func SetPrefix(prefix string) {
	// No-op
}

// Output from standard Go log library. This is a no-op. Provided for compatibility.
func Output(calldepth int, s string) error {
	// No-op (there is no default global logger)
	return nil
}

// ** github.com/Sirupsen/logrus Compatibility **

// Debug from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Debug(args ...interface{}) {
	if l.logger.Level >= logrus.DebugLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Debug()
	}
}

// Debugf from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.logger.Level >= logrus.DebugLevel {
		l.logger.Debugf(format, args...)
	}
}

// Debugln from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Debugln(args ...interface{}) {
	if l.logger.Level >= logrus.DebugLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Debug()
	}
}

// Info from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Info(args ...interface{}) {
	if l.logger.Level >= logrus.InfoLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Info()
	}
}

// Infof from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Infof(format string, args ...interface{}) {
	if l.logger.Level >= logrus.InfoLevel {
		l.logger.Infof(format, args...)
	}
}

// Infoln from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Infoln(args ...interface{}) {
	if l.logger.Level >= logrus.InfoLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Info()
	}
}

// Warn from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Warn(args ...interface{}) {
	if l.logger.Level >= logrus.WarnLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Warn()
	}
}

// Warnf from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.logger.Level >= logrus.WarnLevel {
		l.logger.Warnf(format, args...)
	}
}

// Warnln from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Warnln(args ...interface{}) {
	if l.logger.Level >= logrus.WarnLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Warn()
	}
}

// Warning from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Warning(args ...interface{}) {
	l.Warn(args...)
}

// Warningf from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

// Warningln from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Warningln(args ...interface{}) {
	l.Warnln(args...)
}

// Error from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Error(args ...interface{}) {
	if l.logger.Level >= logrus.ErrorLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Error()
	}
}

// Errorf from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.logger.Level >= logrus.ErrorLevel {
		l.logger.Errorf(format, args...)
	}
}

// Errorln from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) Errorln(args ...interface{}) {
	if l.logger.Level >= logrus.ErrorLevel {
		MarkerMaps.Encode(
			l.logger,
			"",		// event name
			l.name,	// logger name
			map[string]interface{}{
				"args": args,
			},
			map[string]interface{}{},
			nil,
		).Error()
	}
}

// WithField from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) WithField(key string, value interface{}) *logrus.Entry {
	return logrus.NewEntry(l.logger).WithField(key, value)
}

// WithFields from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return logrus.NewEntry(l.logger).WithFields(fields)
}

// WithError from github.com/Sirupsen/logrus library. Provided for compatibility.
func (l *Logger) WithError(err error) *logrus.Entry {
	return logrus.NewEntry(l.logger).WithError(err)
}

// ** Private implementation **

func createLogBuilder(l *logrus.Logger, v logrus.Level, n string) LogBuilder {
	var lb *DefaultLogBuilder = NewDefaultLogBuilder(l, v, n)
	return lb
}
