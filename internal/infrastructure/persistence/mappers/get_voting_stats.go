package mappers

import (
	"encoding/json"

	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

type votingStatsJSON struct {
	TotalProjects      int64 `json:"totalProjects"`
	TotalVotedProjects int64 `json:"totalVotedProjects"`
}

// ToDomainVotingStats converte o JSON retornado por GetVotingStats
// para a entidade de domínio [votacao.VotingStats].
func ToDomainVotingStats(data []byte) (*votacao.VotingStats, error) {
	var row votingStatsJSON
	if err := json.Unmarshal(data, &row); err != nil {
		return nil, err
	}

	return &votacao.VotingStats{
		TotalProjects:      row.TotalProjects,
		TotalVotedProjects: row.TotalVotedProjects,
	}, nil
}
