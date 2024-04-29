package logger

import (
	"log"
	"os"
)

// Глобальный логгер, который можно использовать в любом месте
var GlobalLogger *Logger

// Logger - структура для логирования с уровнями
type Logger struct {
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
}

// NewLogger - создает новый логгер
func NewLogger() *Logger {
	return &Logger{
		infoLog:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		warningLog: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime),
		errorLog:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func init() {
	// Инициализируем глобальный логгер при загрузке пакета
	GlobalLogger = NewLogger()
}

// Методы для записи логов
func (l *Logger) Info(message string) {
	l.infoLog.Println(message)
}

func (l *Logger) Warning(message string) {
	l.warningLog.Println(message)
}

func (l *Logger) Error(message string) {
	l.errorLog.Println(message)
}
