package logger

import "log"

type StdLogger struct {
}

func (s StdLogger) WithField(key string, value interface{}) Logger {
	return s
}

func (s StdLogger) WithFields(fields Fields) Logger {
	return s
}

func (s StdLogger) Trace(args ...interface{}) {
	log.Println(args)
}

func (s StdLogger) Tracef(format string, args ...interface{}) {
	log.Printf(format+"\n", args)
}

func (s StdLogger) Debug(args ...interface{}) {
	log.Println(args)
}

func (s StdLogger) Debugf(format string, args ...interface{}) {
	log.Printf(format+"\n", args)
}

func (s StdLogger) Info(args ...interface{}) {
	log.Println(args)
}

func (s StdLogger) Infof(format string, args ...interface{}) {
	log.Printf(format+"\n", args)
}

func (s StdLogger) Warn(args ...interface{}) {
	log.Println(args)
}

func (s StdLogger) Warnf(format string, args ...interface{}) {
	log.Printf(format+"\n", args)
}

func (s StdLogger) Error(args ...interface{}) {
	log.Println(args)
}

func (s StdLogger) Errorf(format string, args ...interface{}) {
	log.Printf(format+"\n", args)
}

func (s StdLogger) Panic(args ...interface{}) {
	log.Panic(args)
}

func (s StdLogger) Panicf(format string, args ...interface{}) {
	log.Panicf(format+"\n", args)
}

func (s StdLogger) Fatal(args ...interface{}) {
	log.Fatal(args)
}

func (s StdLogger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format+"\n", args)
}
