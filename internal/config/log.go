package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// LogLevel representa o nivel de severidade de uma mensagem de log
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

type Logger struct {
	writer  io.Writer
	mu      sync.RWMutex
	level   LogLevel
	encoder *json.Encoder
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "") // Garante que nao haja indentacao para melhor performance

	return &Logger{
		writer:  writer,
		level:   INFO,
		encoder: encoder,
	}
}

// Garante que o nivel de severidade da mensagem de log seja o nivel informado
func (l *Logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// log cria uma entrada de log estruturada
func (l *Logger) log(level LogLevel, msg string, fields map[string]interface{}) {
	if level < l.level {
		return
	}

	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level.String(),
		Message:   msg,
		Fields:    fields,
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if err := l.encoder.Encode(entry); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding log entry: %v\n", err)
	}
}

// String retorna a representacao em string de um nivel de log
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Garante que uma mensagem de log de nivel DEBUG seja registrada
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	mergedFields := mergeFields(fields...)
	l.log(DEBUG, msg, mergedFields)
}

// Garante que uma mensagem de log de nivel INFO seja registrada
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	mergedFields := mergeFields(fields...)
	l.log(INFO, msg, mergedFields)
}

// Garante que uma mensagem de log de nivel WARN seja registrada
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	mergedFields := mergeFields(fields...)
	l.log(WARN, msg, mergedFields)
}

// Garante que uma mensagem de log de nivel ERROR seja registrada
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	mergedFields := mergeFields(fields...)
	l.log(ERROR, msg, mergedFields)
}

// Garante que varios mapas de campos sejam combinados em um unico mapa
func mergeFields(fields ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, f := range fields {
		for k, v := range f {
			result[k] = v
		}
	}
	return result
}
