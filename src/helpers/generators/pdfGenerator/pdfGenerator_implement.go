package pdfGenerator

import (
	"log"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wk struct {
	rootPath string
}

func NewWkHtmlToPdf(rootPath string) PDFGeneratorInterface {
	return &wk{rootPath: rootPath}
}

func (w *wk) Create(htmlFile string) (string, error) {
	f, err := os.Open(htmlFile)
	if err != nil {
		return "", err

	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	page := wkhtmltopdf.NewPage(w.rootPath)

	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	if err := pdfg.Create(); err != nil {
		return "", err
	}

	fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

	if err := pdfg.WriteFile(fileName); err != nil {
		return "", err
	}

	return "", nil
}
