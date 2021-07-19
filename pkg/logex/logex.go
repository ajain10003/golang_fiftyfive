package logex

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SetupLogging configures consistent logging output in a standard format for.
func SetupLogging() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

//DisableLogging stops the logger. Useful for quieting tests.
func DisableLogging() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
}

// SetupAndBuild sets up and builds logger
func SetupAndBuild(appName string) (*zap.Logger, func()) {
	SetupLogging()
	l := Build(appName)
	f := func() {
		//deliberately ignoring error from sync call. If there's a problem, it's too late now!
		_ = l.Sync()
	}
	return l, f
}

// Build builds the default zap logger, and sets the global zap logger to the configured logger instance.
func Build(appName string) *zap.Logger {
	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(time.Now().String())
			},
			CallerKey:      "caller",
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("err making logger: %+v", err)
	}

	logger = logger.With(zap.String("app", appName))
	_ = zap.ReplaceGlobals(logger)

	return logger
}
