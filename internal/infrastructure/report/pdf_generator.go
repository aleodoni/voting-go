package report

import (
	"bytes"
	"context"
	"fmt"

	"github.com/go-pdf/fpdf"

	domainRelatorio "github.com/aleodoni/voting-go/internal/domain/relatorio"
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

// PDFRelatorioReuniaoGenerator implementa [domainRelatorio.Generator] gerando um PDF.
type PDFRelatorioReuniaoGenerator struct{}

func NewPDFRelatorioReuniaoGenerator() *PDFRelatorioReuniaoGenerator {
	return &PDFRelatorioReuniaoGenerator{}
}

func (g *PDFRelatorioReuniaoGenerator) Gera(_ context.Context, data domainRelatorio.ReuniaoOutput) ([]byte, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(marginL, marginT, marginR)
	pdf.SetAutoPageBreak(true, 15)

	pdf.AddUTF8FontFromBytes("DejaVu", "", dejaVuSans)
	pdf.AddUTF8FontFromBytes("DejaVu", "B", dejaVuSansBold)

	pdf.RegisterImageOptionsReader(
		"logo",
		fpdf.ImageOptions{ImageType: "PNG"},
		bytes.NewReader(logoPNG),
	)

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

func drawHeader(pdf *fpdf.Fpdf, data domainRelatorio.ReuniaoOutput) {
	pdf.ImageOptions("logo", marginL, marginT, 18, 18, false, fpdf.ImageOptions{ImageType: "PNG"}, 0, "")

	pdf.SetXY(marginL+21, marginT+1)
	pdf.SetFont("DejaVu", "B", 16)
	pdf.SetTextColor(0, 80, 160)
	pdf.Cell(contentW-21, 8, "Câmara Municipal de Curitiba")

	pdf.SetXY(marginL, marginT+11)
	pdf.SetFont("DejaVu", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(contentW, 7, "Relatório de Votação", "", 1, "C", false, 0, "")

	pdf.Ln(2)
	drawLine(pdf)
	pdf.Ln(3)

	pdf.SetFont("DejaVu", "B", 11)
	pdf.Cell(contentW, rowH, data.ConDesc)
	pdf.Ln(rowH)

	pdf.SetFont("DejaVu", "B", 10)
	pdf.Cell(contentW, smallH, fmt.Sprintf("Número da reunião: %s", data.RecNumero))
	pdf.Ln(smallH)

	pdf.Cell(contentW, smallH, fmt.Sprintf("Reunião: %s", data.RecTipoReuniao))
	pdf.Ln(smallH)

	pdf.Cell(contentW, smallH, fmt.Sprintf("Data: %s", data.RecData))
	pdf.Ln(smallH + 2)

	drawLine(pdf)
	pdf.Ln(4)
}

func drawProjeto(pdf *fpdf.Fpdf, p domainRelatorio.ProjetoItem) {
	pdf.SetFont("DejaVu", "B", 12)
	pdf.Cell(contentW, rowH, p.CodigoProposicao)
	pdf.Ln(rowH + 1)

	pdf.SetFont("DejaVu", "B", 10)
	pdf.Cell(colLabelW, smallH, "Iniciativa")
	pdf.SetFont("DejaVu", "", 10)
	pdf.Cell(colValueW, smallH, p.Iniciativa)
	pdf.Ln(smallH)

	pdf.SetFont("DejaVu", "B", 10)
	pdf.Cell(colLabelW, smallH, "Relator")
	pdf.SetFont("DejaVu", "", 10)
	pdf.Cell(colValueW, smallH, p.Relator)
	pdf.Ln(smallH + 1)

	drawLine(pdf)
	pdf.Ln(3)

	if p.Votacao == nil {
		pdf.SetFont("DejaVu", "", 10)
		pdf.Cell(contentW, smallH, "Sem votação registrada")
		pdf.Ln(smallH)
		return
	}

	pdf.SetFont("DejaVu", "", 10)
	pdf.Cell(contentW, smallH, "Como votaram os vereadores")
	pdf.Ln(smallH + 1)

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

	pdf.SetFont("DejaVu", "", 10)
	for _, t := range totaisOrdem {
		drawLine(pdf)
		total := p.Votacao.Totais[t.key]
		pdf.SetX(marginL)
		pdf.Cell(contentW-10, smallH, t.label)
		pdf.CellFormat(10, smallH, fmt.Sprintf("%d", total), "", 1, "R", false, 0, "")
	}

	if len(p.Votacao.Votos) == 0 {
		return
	}

	pdf.Ln(4)
	pdf.SetFont("DejaVu", "B", 10)
	pdf.Cell(contentW, smallH, "Votos")
	pdf.Ln(smallH + 1)

	pdf.SetFont("DejaVu", "", 10)
	for _, voto := range p.Votacao.Votos {
		drawLine(pdf)
		pdf.SetX(marginL)
		pdf.Cell(contentW-30, smallH, voto.UsuarioNome)
		pdf.CellFormat(30, smallH, opcaoVotoDesc(voto.Opcao), "", 1, "R", false, 0, "")

		if voto.Restricao != "" {
			pdf.SetX(marginL + 4)
			pdf.SetFont("DejaVu", "", 9)
			pdf.Cell(contentW-4, smallH, fmt.Sprintf("Restrição: %s", voto.Restricao))
			pdf.Ln(smallH)
			pdf.SetFont("DejaVu", "", 10)
		}

		if voto.VotoContrario != "" {
			pdf.SetX(marginL + 4)
			pdf.SetFont("DejaVu", "", 9)
			pdf.Cell(contentW-4, smallH, fmt.Sprintf("Voto contrário: %s", voto.VotoContrario))
			pdf.Ln(smallH)
			pdf.SetFont("DejaVu", "", 10)
		}
	}
}

func drawLine(pdf *fpdf.Fpdf) {
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(marginL, pdf.GetY(), marginL+contentW, pdf.GetY())
	pdf.SetDrawColor(0, 0, 0)
}

func opcaoVotoDesc(o string) string {
	switch o {
	case "F":
		return "Favorável"
	case "R":
		return "Contrário"
	case "C":
		return "Com restrição"
	case "V":
		return "Vistas"
	case "A":
		return "Abstenção"
	default:
		return o
	}
}
