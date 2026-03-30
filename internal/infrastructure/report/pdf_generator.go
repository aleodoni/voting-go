package report

import (
	"bytes"
	"context"
	"fmt"

	domainRelatorio "github.com/aleodoni/voting-go/internal/domain/relatorio"
	"github.com/go-pdf/fpdf"
)

const (
	pageW     = 210.0
	marginL   = 15.0
	marginR   = 15.0
	marginT   = 15.0
	contentW  = pageW - marginL - marginR
	colLabelW = 40.0
	colValueW = contentW - colLabelW
	rowH      = 7.0
	smallH    = 6.0
)

// PDFRelatorioReuniaoGenerator gera um PDF profissional de relatório
type PDFRelatorioReuniaoGenerator struct{}

func NewPDFRelatorioReuniaoGenerator() *PDFRelatorioReuniaoGenerator {
	return &PDFRelatorioReuniaoGenerator{}
}

func (g *PDFRelatorioReuniaoGenerator) Gera(_ context.Context, data domainRelatorio.ReuniaoOutput) ([]byte, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(marginL, marginT, marginR)
	pdf.SetAutoPageBreak(true, 25)

	// Fontes
	pdf.AddUTF8FontFromBytes("DejaVu", "", dejaVuSans)
	pdf.AddUTF8FontFromBytes("DejaVu", "B", dejaVuSansBold)

	// Logo
	pdf.RegisterImageOptionsReader(
		"logo",
		fpdf.ImageOptions{ImageType: "PNG"},
		bytes.NewReader(logoPNG),
	)

	pdf.AliasNbPages("")

	// Footer automático
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("DejaVu", "", 9)
		pdf.SetTextColor(100, 100, 100)
		pdf.CellFormat(0, smallH, fmt.Sprintf("Página %d de {nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})

	// Geração de páginas
	for _, projeto := range data.Projetos {
		pdf.AddPage()
		drawHeader(pdf, data)
		drawProjeto(pdf, projeto)
	}

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, fmt.Errorf("PDFRelatorioReuniaoGenerator: %w", err)
	}
	return buf.Bytes(), nil
}

// --- HEADER ---
func drawHeader(pdf *fpdf.Fpdf, data domainRelatorio.ReuniaoOutput) {
	// Logo
	pdf.ImageOptions("logo", marginL, marginT, 18, 18, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")

	// Título
	pdf.SetXY(marginL+21, marginT+1)
	pdf.SetFont("DejaVu", "B", 16)
	pdf.SetTextColor(0, 80, 160)
	pdf.CellFormat(contentW-21, 8, "Câmara Municipal de Curitiba", "", 0, "L", false, 0, "")

	// Subtítulo
	pdf.SetXY(marginL+21, marginT+11)
	pdf.SetFont("DejaVu", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(contentW-21, 6, "Relatório de Votação", "", 0, "L", false, 0, "")

	pdf.Ln(20)

	// Descrição da comissão
	pdf.SetFont("DejaVu", "B", 11)
	pdf.MultiCell(contentW, rowH, data.ConDesc, "", "L", false)
	pdf.Ln(2)

	// Informações da reunião (linha única, quebrando se necessário)
	pdf.SetFont("DejaVu", "B", 10)
	pdf.MultiCell(contentW, smallH, fmt.Sprintf("Número: %s | Tipo: %s | Data: %s", data.RecNumero, data.RecTipoReuniao, data.RecData), "", "L", false)
	pdf.Ln(2)

	drawLine(pdf)
	pdf.Ln(4)
}

// --- PROJETOS ---
func drawProjeto(pdf *fpdf.Fpdf, p domainRelatorio.ProjetoItem) {
	// Cabeçalho do projeto com fundo suave
	pdf.SetFillColor(220, 230, 250) // azul claro suave
	pdf.SetFont("DejaVu", "B", 12)
	pdf.MultiCell(contentW, rowH, p.CodigoProposicao, "", "L", true)
	pdf.Ln(1)

	// Iniciativa e Relator
	pdf.SetFont("DejaVu", "B", 10)
	pdf.CellFormat(colLabelW, smallH, "Iniciativa", "", 0, "L", false, 0, "")
	pdf.SetFont("DejaVu", "", 10)
	pdf.MultiCell(colValueW, smallH, p.Iniciativa, "", "L", false)

	pdf.SetFont("DejaVu", "B", 10)
	pdf.CellFormat(colLabelW, smallH, "Relator", "", 0, "L", false, 0, "")
	pdf.SetFont("DejaVu", "", 10)
	pdf.MultiCell(colValueW, smallH, p.Relator, "", "L", false)

	drawLine(pdf)
	pdf.Ln(3)

	if p.Votacao == nil {
		pdf.SetFont("DejaVu", "", 10)
		pdf.MultiCell(contentW, smallH, "Sem votação registrada", "", "L", false)
		return
	}

	// Votação
	pdf.SetFont("DejaVu", "B", 10)
	pdf.MultiCell(contentW, smallH, "Como votaram os vereadores", "", "L", false)
	pdf.Ln(1)

	// Alterna cores de linha para votos
	rowColors := []struct{ r, g, b int }{{245, 245, 245}, {255, 255, 255}}
	colorIndex := 0

	for _, voto := range p.Votacao.Votos {
		pdf.SetFillColor(rowColors[colorIndex].r, rowColors[colorIndex].g, rowColors[colorIndex].b)
		colorIndex = 1 - colorIndex

		drawLine(pdf)
		pdf.SetX(marginL)
		pdf.SetFont("DejaVu", "", 10)
		pdf.CellFormat(contentW-30, smallH, voto.UsuarioNome, "", 0, "L", true, 0, "")
		pdf.CellFormat(30, smallH, opcaoVotoDesc(voto.Opcao), "", 1, "R", true, 0, "")

		// Restrição
		if voto.Restricao != "" {
			pdf.SetX(marginL + 4)
			pdf.SetFont("DejaVu", "", 9)
			pdf.SetTextColor(80, 80, 80)
			pdf.MultiCell(contentW-4, smallH, fmt.Sprintf("Restrição: %s", voto.Restricao), "", "L", false)
			pdf.SetTextColor(0, 0, 0)
		}

		// Motivo do voto contrário
		if voto.VotoContrario != "" {
			pdf.SetX(marginL + 4)
			pdf.SetFont("DejaVu", "", 9)
			pdf.SetTextColor(80, 80, 80)
			pdf.MultiCell(contentW-4, smallH, fmt.Sprintf("Motivo: %s", voto.VotoContrario), "", "L", false)
			pdf.SetTextColor(0, 0, 0)
		}
	}

	pdf.Ln(2)

	// Totais
	totaisOrdem := []struct {
		key   string
		label string
	}{
		{"F", "TOTAL VOTOS FAVORÁVEL"},
		{"C", "TOTAL VOTOS FAVORÁVEL COM RESTRIÇÕES"},
		{"R", "TOTAL VOTOS CONTRÁRIO"},
		{"V", "TOTAL VOTOS VISTAS"},
		{"A", "TOTAL VOTOS ABSTENÇÃO"},
	}

	pdf.SetFont("DejaVu", "B", 10)
	for _, t := range totaisOrdem {
		drawLine(pdf)
		total := p.Votacao.Totais[t.key]
		pdf.SetX(marginL)
		pdf.CellFormat(contentW-10, smallH, t.label, "", 0, "L", false, 0, "")
		pdf.CellFormat(10, smallH, fmt.Sprintf("%d", total), "", 1, "R", false, 0, "")
	}
}

// --- LINHA ---
func drawLine(pdf *fpdf.Fpdf) {
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(marginL, pdf.GetY(), marginL+contentW, pdf.GetY())
	pdf.SetDrawColor(0, 0, 0)
}

// --- OPÇÃO VOTO ---
func opcaoVotoDesc(o string) string {
	switch o {
	case "F":
		return "FAVORÁVEL"
	case "R":
		return "CONTRÁRIO"
	case "C":
		return "FAVORÁVEL COM RESTRIÇÕES"
	case "V":
		return "VISTAS"
	case "A":
		return "ABSTENÇÃO"
	default:
		return o
	}
}
