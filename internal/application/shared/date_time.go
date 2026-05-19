package shared

import "time"

// GetCurrentDate retorna a data atual sem componente de horário (00:00:00),
// preservando o fuso horário local do sistema.
func GetCurrentDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}
