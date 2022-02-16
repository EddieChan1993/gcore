package glog

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	initStderr()
	ResetToProduction()
}

type Level int

const (
	DebugLevel Level = iota + 1
	InfoLevel
	ErrorLevel
)

type Format int

const (
	Json Format = iota + 1
	Human
)

type Receiver int

const (
	Console Receiver = iota + 1
	File
)

func ResetToProduction() {
	Reset(InfoLevel, Json, File, fileName)
}

func ResetToDevelopment() {
	Reset(DebugLevel, Human, Console, fileName)
}

func ResetLevel(logLevel Level) {
	Reset(logLevel, format, receiver, fileName)
}
func ResetFormat(logFormat Format) {
	Reset(level, logFormat, receiver, fileName)
}
func ResetReceiver(logReceiver Receiver) {
	Reset(level, format, logReceiver, fileName)
}
func ResetFileName(logFileName string) {
	Reset(level, format, receiver, logFileName)
}

func Reset(logLevel Level, logFormat Format, logReceiver Receiver, logFileName string) error {
	if logFileName == "" {
		exeName := strings.Split(os.Args[0], "/")
		logFileName = exeName[len(exeName)-1]
	}

	encoder, err := createEncoder(logFormat)
	if err != nil {
		return err
	}

	core, err := createWriteSyncer(encoder, logReceiver, logFileName, logLevel)
	if err != nil {
		return err
	}

	level = logLevel
	format = logFormat
	receiver = logReceiver
	fileName = logFileName

	callerOption := zap.AddCaller()
	callerSkip := zap.AddCallerSkip(1)

	logger = zap.New(core, callerOption, callerSkip).Sugar()
	return nil
}

func Debugw(msg string, keysAndValues ...interface{}) {
	defer logger.Sync()
	logger.Debugw(msg, keysAndValues...)
}

func Debugr(msg, requestID string, keysAndValues ...interface{}) {
	defer logger.Sync()
	t := []interface{}{"request_id", requestID}
	t = append(t, keysAndValues...)
	logger.Debugw(msg, t)
}

func Infow(msg string, keysAndValues ...interface{}) {
	defer logger.Sync()
	logger.Infow(msg, keysAndValues...)
}

func Infor(msg, requestID string, keysAndValues ...interface{}) {
	defer logger.Sync()
	t := []interface{}{"request_id", requestID}
	t = append(t, keysAndValues...)
	logger.Infow(msg, t...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	defer logger.Sync()
	logger.Warnw(msg, keysAndValues...)
}

func Warnr(msg, requestID string, keysAndValues ...interface{}) {
	defer logger.Sync()
	t := []interface{}{"request_id", requestID}
	t = append(t, keysAndValues...)
	logger.Warnw(msg, t...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	defer logger.Sync()
	logger.Errorw(msg, keysAndValues...)
}

func Errorr(msg, requestID string, keysAndValues ...interface{}) {
	defer logger.Sync()
	t := []interface{}{"request_id", requestID}
	t = append(t, keysAndValues...)
	logger.Errorw(msg, t...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	defer logger.Sync()
	logger.Fatalw(msg, keysAndValues...)
}

func Fatalr(msg, requestID string, keysAndValues ...interface{}) {
	defer logger.Sync()
	t := []interface{}{"request_id", requestID}
	t = append(t, keysAndValues...)
	logger.Fatalw(msg, t...)
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	defer logger.Sync()
	logger.Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	defer logger.Sync()
	logger.Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	defer logger.Sync()
	logger.Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	defer logger.Sync()
	logger.Error(args)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	defer logger.Sync()
	logger.Panic(args)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	defer logger.Sync()
	logger.Fatal(args)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	defer logger.Sync()
	logger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	defer logger.Sync()
	logger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	defer logger.Sync()
	logger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	defer logger.Sync()
	logger.Errorf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message.
func Fatalf(template string, args ...interface{}) {
	defer logger.Sync()
	logger.Fatalf(template, args...)
}

func createEncoder(format Format) (zapcore.Encoder, error) {
	var encoder zapcore.Encoder
	switch format {
	case Json:
		encoderConfig := ecszap.NewDefaultEncoderConfig().ToZapCoreEncoderConfig()
		encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case Human:
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000Z0700")
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderConfig.ConsoleSeparator = " | "
		encoderConfig.FunctionKey = "func"
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		return nil, fmt.Errorf("unkown format %v", format)
	}
	return encoder, nil
}

func createWriteSyncer(encoder zapcore.Encoder, receiver Receiver, fileName string, logLevel Level) (zapcore.Core, error) {
	var zapLevel zapcore.Level
	switch logLevel {
	case DebugLevel:
		zapLevel = zapcore.DebugLevel
	case InfoLevel:
		zapLevel = zapcore.InfoLevel
	case ErrorLevel:
		zapLevel = zapcore.ErrorLevel
	default:
		return nil, fmt.Errorf("unkown loglevel:%v", logLevel)
	}

	var writeSyncer zapcore.WriteSyncer
	switch receiver {
	case Console:
		writeSyncer = zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(encoder, writeSyncer, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapLevel
		}))
		return core, nil
	case File:
		// 实现两个判断日志等级的interface (其实 zapcore.*Level 自身就是 interface)
		infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapLevel
		})
		warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.WarnLevel
		})

		// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
		infoWriter := getWriter("./log/" + fileName + ".log.json")
		warnWriter := stdErrFileHandler

		// 最后创建具体的Logger
		core := zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
		)
		return core, nil
	default:
		return nil, fmt.Errorf("unkown reciver %v", receiver)
	}
}

func initStderr() {
	if err := os.MkdirAll("./log", os.ModePerm); err != nil {
		panic(err)
	}

	if runtime.GOOS == "windows" {
		return
	}

	exeName := strings.Split(os.Args[0], "/")
	fileName := exeName[len(exeName)-1]

	stdErrFile := "./log/" + fileName + ".error.json"
	if runtime.GOOS == "windows" {
		return
	}

	file, err := os.OpenFile(stdErrFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	stdErrFileHandler = file
	if err := syscall.Dup2(int(stdErrFileHandler.Fd()), int(os.Stderr.Fd())); err != nil {
		Fatalw("failed to init stderr", "err", err)
		panic(err)
	}

	// 内存回收前关闭文件描述符
	runtime.SetFinalizer(stdErrFileHandler, func(fd *os.File) {
		fd.Close()
	})
}

func getWriter(filename string) io.Writer {
	// 生成 rotatelogs 的Logger 实际生成的文件名 xxx.log.YYmmddHH
	// xxx.log是指向最新日志的链接
	// 保存30天内的日志，每24小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

var stdErrFileHandler *os.File
var logger *zap.SugaredLogger
var level Level
var format Format
var receiver Receiver
var fileName string
