package logger

import "github.com/kordar/nazalog"

type nazalogAdapt struct {
	l nazalog.Logger
}

func NewNazalogAdapt() Logger {
	return nazalogAdapt{l: nazalog.GetGlobalLogger()}
}

func (n nazalogAdapt) WithField(key string, value interface{}) Logger {
	//TODO implement me
	panic("implement me")
}

func (n nazalogAdapt) WithFields(fields Fields) Logger {
	//TODO implement me
	panic("implement me")
}

func (n nazalogAdapt) Trace(args ...interface{}) {
	//TODO implement me
	n.l.Trace(args...)
}

func (n nazalogAdapt) Tracef(format string, args ...interface{}) {
	//TODO implement me
	n.l.Tracef(format, args...)
}

func (n nazalogAdapt) Debug(args ...interface{}) {
	//TODO implement me
	n.l.Debug(args...)
}

func (n nazalogAdapt) Debugf(format string, args ...interface{}) {
	//TODO implement me
	n.l.Debugf(format, args...)
}

func (n nazalogAdapt) Info(args ...interface{}) {
	//TODO implement me
	n.l.Info(args...)
}

func (n nazalogAdapt) Infof(format string, args ...interface{}) {
	//TODO implement me
	n.l.Infof(format, args...)
}

func (n nazalogAdapt) Warn(args ...interface{}) {
	//TODO implement me
	n.l.Warn(args...)
}

func (n nazalogAdapt) Warnf(format string, args ...interface{}) {
	//TODO implement me
	n.l.Warnf(format, args...)
}

func (n nazalogAdapt) Error(args ...interface{}) {
	//TODO implement me
	n.l.Error(args...)
}

func (n nazalogAdapt) Errorf(format string, args ...interface{}) {
	//TODO implement me
	n.l.Errorf(format, args...)
}

func (n nazalogAdapt) Panic(args ...interface{}) {
	//TODO implement me
	n.l.Panic(args...)
}

func (n nazalogAdapt) Panicf(format string, args ...interface{}) {
	//TODO implement me
	n.l.Panicf(format, args...)
}

func (n nazalogAdapt) Fatal(args ...interface{}) {
	//TODO implement me
	n.l.Fatal(args...)
}

func (n nazalogAdapt) Fatalf(format string, args ...interface{}) {
	//TODO implement me
	n.l.Fatalf(format, args...)
}
