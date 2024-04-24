package config

type LoggerConfig interface {
	Filename() string
	ServiceName() string
	MaxSize() int
	MaxBackups() int
	MaxAge() int
	Level() string
}

type loggerConfig struct {
	serviceName string
	filename    string
	maxSize     int
	maxBackups  int
	maxAge      int
	level       string
}

func (l *loggerConfig) ServiceName() string {
	return l.serviceName
}

func (l *loggerConfig) Filename() string {
	return l.filename
}

func (l *loggerConfig) MaxSize() int {
	return l.maxSize
}

func (l *loggerConfig) MaxBackups() int {
	return l.maxBackups
}

func (l *loggerConfig) MaxAge() int {
	return l.maxAge
}

func (l *loggerConfig) Level() string {
	return l.level
}

func NewLoggerConfig() (LoggerConfig, error) {
	const maxAge = 7
	const maxBackups = 3
	const maxSize = 10
	return &loggerConfig{
		filename:    "logs/pastebin-log.json",
		maxSize:     maxSize, // megabytes
		maxBackups:  maxBackups,
		maxAge:      maxAge, // days
		level:       "DEBUG",
		serviceName: "GoPastebin App",
	}, nil
}
