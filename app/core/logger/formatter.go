package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Formatter struct{}

const TimeFormat = "2006-01-02 15:04:05"

func (s *Formatter) Format(entry *logrus.Entry) ([]byte, error) {

	msg := fmt.Sprintf("[%s] [%s] %s\n", time.Now().Local().Format(TimeFormat), strings.ToUpper(entry.Level.String()), entry.Message)

	return []byte(msg), nil
}
