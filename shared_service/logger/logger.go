package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type Logger interface {
	Error(v ...any)
	Info(v ...any)
	Debug(v ...any)
	Trace(v ...any)
	Fatal(v ...any)
	Panic(v ...any)
	Print(v ...any)
	Println(v ...any)
}

type logger struct {
	prefixError []byte
	prefixInfo  []byte
	prefixDebug []byte
	prefixTrace []byte
	prefixFatal []byte
	prefixPanic []byte
	writer      io.Writer
	pool        *Pool[[]byte] // Пул строк для логирования
}

type Pool[T any] struct {
	pool sync.Pool
}

func NewPool[T any]() *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{
			New: func() any {
				return new(T)
			},
		},
	}
}

func (p *Pool[T]) Get() *T {
	return p.pool.Get().(*T)
}

func (p *Pool[T]) Put(item *T) {
	p.pool.Put(item)
}

func NewWithWriter(out io.Writer) Logger {

	color := func(c, prefix string) []byte {
		if os.Stdout == out {
			return []byte(fmt.Sprintf("%s%s: %s", c, prefix, colorReset))
		}
		return []byte(fmt.Sprintf("%s: ", prefix))
	}

	return &logger{
		prefixError: color(colorRed, "ERROR"),
		prefixInfo:  color(colorGreen, "INFO"),
		prefixDebug: color(colorYellow, "DEBUG"),
		prefixTrace: color(colorBlue, "TRACE"),
		prefixFatal: color(colorPurple, "FATAL"),
		prefixPanic: color(colorCyan, "PANIC"),
		writer:      out,
		pool:        NewPool[[]byte](),
	}
}

func New() Logger {
	return NewWithWriter(os.Stdout)
}

func (l *logger) logWithCaller(prefix []byte, v ...any) {
	// Получаем информацию о файле и строке
	_, file, line, ok := runtime.Caller(2)
	if ok {
		// Форматируем строку с файлом, строкой и временем
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		prefixFile := fmt.Sprintf("%s[%s] %s:%d ", timestamp, prefix, file, line)
		// Получаем строку из пула
		buf := l.pool.Get()
		// Собираем сообщение
		*buf = append(*buf, []byte(prefixFile)...)
		*buf = append(*buf, fmt.Sprint(v...)...)
		// Печатаем сообщение
		fmt.Fprintln(l.writer, string(*buf))
		// Возвращаем строку в пул
		l.pool.Put(buf)
	} else {
		// Если не удалось получить информацию о файле/строке
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		// Получаем строку из пула
		buf := l.pool.Get()
		// Собираем сообщение
		*buf = append(*buf, []byte(fmt.Sprintf("%s[%s] ", timestamp, prefix))...)
		*buf = append(*buf, fmt.Sprint(v...)...)
		// Печатаем сообщение
		fmt.Fprintln(l.writer, string(*buf))
		// Возвращаем строку в пул
		l.pool.Put(buf)
	}
}

func (l *logger) Error(v ...any) {
	l.logWithCaller(l.prefixError, v...)
}

func (l *logger) Info(v ...any) {
	l.logWithCaller(l.prefixInfo, v...)
}

func (l *logger) Debug(v ...any) {
	l.logWithCaller(l.prefixDebug, v...)
}

func (l *logger) Trace(v ...any) {
	l.logWithCaller(l.prefixTrace, v...)
}

func (l *logger) Fatal(v ...any) {
	l.logWithCaller(l.prefixFatal, v...)
	os.Exit(1)
}

func (l *logger) Panic(v ...any) {
	l.logWithCaller(l.prefixPanic, v...)
	panic(fmt.Sprint(v...))
}

func (l *logger) Print(v ...any) {
	l.logWithCaller([]byte{}, v...)
}

func (l *logger) Println(v ...any) {
	l.logWithCaller([]byte{}, v...)
}
