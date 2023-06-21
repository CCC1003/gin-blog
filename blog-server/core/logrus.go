package core

import (
	"Blog/global"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

type fileDateHook struct {
	file     *os.File
	logPath  string
	fileDate string
	appName  string
}

func (hook *fileDateHook) Fire(entry *logrus.Entry) error {
	//TODO implement me
	//时间相同
	timer := time.Now().Format("2006-01-02-15-04")
	line, _ := entry.String()
	if timer == hook.fileDate {
		hook.file.Write([]byte(line))
	}
	//时间不同
	hook.file.Close()
	//创建新的目录文件夹
	os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, timer), os.ModePerm)
	fileName := fmt.Sprintf("%s/%s/%s.log", hook.logPath, timer, hook.appName)

	//读取新的日志文件
	hook.file, _ = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

func (hook *fileDateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

type LogFormatter struct{}

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	logger := global.Config.Logger
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	//对象是否设置了logrus.SetReportCaller(true)  显示文件名,行号和函数等
	if entry.HasCaller() {
		//显示函数
		funcVal := entry.Caller.Function
		//显示文件名和行号
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", logger.Prefix, timestamp,
			levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", logger.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}
func InitLogger(logPath, appName string) *logrus.Logger {
	mLog := logrus.New()               //新建一个日志实例
	mLog.SetReportCaller(true)         //设置是否开启返回文件路径和行号等信息
	mLog.SetFormatter(&LogFormatter{}) //设置自定义的日志格式

	fileDate := time.Now().Format("2006-01-02-15-04")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	fileName := fmt.Sprintf("%s/%s/%s.log", logPath, fileDate, appName)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	writer := []io.Writer{
		file,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writer...) //同时些文件和屏幕
	mLog.SetOutput(fileAndStdoutWriter)              //设置输出类型
	if err != nil {
		logrus.Error(err)
		return nil
	}
	fileHook := fileDateHook{file, logPath, fileDate, fileName}
	logrus.AddHook(&fileHook)
	level, err := logrus.ParseLevel(global.Config.Logger.Level) //接收logrus中对应日志等级
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) //设置日志等级
	//InitDefaultLogger()
	return mLog
}
func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
