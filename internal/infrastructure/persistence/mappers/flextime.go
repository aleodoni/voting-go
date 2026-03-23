package mappers

import (
	"fmt"
	"strings"
	"time"
)

type flexTime struct {
	time.Time
}

func (ft *flexTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)

	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		ft.Time = t
		return nil
	}

	formats := []string{
		"2006-01-02T15:04:05.999999",
		"2006-01-02T15:04:05",
	}
	for _, f := range formats {
		if t, err := time.Parse(f, s); err == nil {
			ft.Time = t
			return nil
		}
	}

	return fmt.Errorf("não foi possível parsear data: %s", s)
}
