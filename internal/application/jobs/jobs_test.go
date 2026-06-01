package jobs_test

import (
	"context"
	"errors"
	"testing"

	ucJobs "github.com/aleodoni/voting-go/internal/application/jobs"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

// ─────────────────────────────────────────────────────────────────────────────
// ExecutaSincroniaJobUseCase
// ─────────────────────────────────────────────────────────────────────────────

func TestExecutaSincroniaJob_Sucesso(t *testing.T) {
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	uc := ucJobs.NewExecutaSincroniaJobUseCase(sincroniaRepo)
	result, err := uc.Execute(context.Background())

	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
	if result == nil {
		t.Fatal("esperava sincronia retornada, obteve nil")
	}
	if sincroniaRepo.SyncCalls != 1 {
		t.Fatalf("esperava 1 chamada a Sync, obteve %d", sincroniaRepo.SyncCalls)
	}
}

func TestExecutaSincroniaJob_ErroRepositorio(t *testing.T) {
	sincroniaRepo := fakes.NewFakeSincroniaRepository()
	erroEsperado := errors.New("falha na sincronização")
	sincroniaRepo.SyncErr = erroEsperado

	uc := ucJobs.NewExecutaSincroniaJobUseCase(sincroniaRepo)
	_, err := uc.Execute(context.Background())

	if !errors.Is(err, erroEsperado) {
		t.Fatalf("esperava erro do repositório, obteve: %v", err)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// FechaVotacoesAbertasJobUseCase
// ─────────────────────────────────────────────────────────────────────────────

func TestFechaVotacoesAbertasJob_Sucesso(t *testing.T) {
	jobRepo := fakes.NewFakeJobRepository()

	uc := ucJobs.NewFechaVotacoesAbertasJobUseCase(jobRepo)
	err := uc.Execute(context.Background())

	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
	if jobRepo.FecharVotacoesAbertasCalls != 1 {
		t.Fatalf("esperava 1 chamada a FecharVotacoesAbertas, obteve %d", jobRepo.FecharVotacoesAbertasCalls)
	}
}

func TestFechaVotacoesAbertasJob_ErroRepositorio(t *testing.T) {
	jobRepo := fakes.NewFakeJobRepository()
	erroEsperado := errors.New("falha ao fechar votações")
	jobRepo.FecharVotacoesAbertasErr = erroEsperado

	uc := ucJobs.NewFechaVotacoesAbertasJobUseCase(jobRepo)
	err := uc.Execute(context.Background())

	if !errors.Is(err, erroEsperado) {
		t.Fatalf("esperava erro do repositório, obteve: %v", err)
	}
}
