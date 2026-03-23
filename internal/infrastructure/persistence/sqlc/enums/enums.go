// Package enums defines the enum types used in the application.
package enums

type StatusVotacao string

const (
	StatusVotacaoA StatusVotacao = "A" // Aberta
	StatusVotacaoF StatusVotacao = "F" // Fechada
	StatusVotacaoV StatusVotacao = "V" // Votando
	StatusVotacaoC StatusVotacao = "C" // Cancelada
)

type OpcaoVoto string

const (
	OpcaoVotoF OpcaoVoto = "F" // Favorável
	OpcaoVotoR OpcaoVoto = "R" // Rejeição
	OpcaoVotoC OpcaoVoto = "C" // Contrário
	OpcaoVotoV OpcaoVoto = "V" // Vistas
	OpcaoVotoA OpcaoVoto = "A" // Abstenção
)
