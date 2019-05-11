package presenter

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

// The presentation to be turned into
// a sequence of PDF pages
type Presentation interface {

	// Will put the given text on each an
	// every screen
	HeaderText(string)

	// Expects an image as base64 encoded data
	HeaderPNGbase64(string)

	// Will put the given text on each an
	// every screen
	FooterText(string)

	// Adds a screen to the presentation
	AddScreen(Screen)

	// Sets the path of the output file to be written
	OutputFilePath(string)

	// Actually creates the PDF file at the
	// given path
	Render()
}

// A single screen of the PDF presentation
type Screen interface {

	// The headline of a given Screen
	Headline(string)

	// The main text of a given Screen
	// expects plain text (optional)
	Text(string)

	// A PNG image to be shown on the
	// screen (optional)
	PNGbase64(string)
}

// implementation

type presentation struct {
	header     *header
	footer     *footer
	outputPath string
	screens    []Screen
}

func (p *presentation) HeaderPNGbase64(png string) {
	p.header.png = png
}

func (p *presentation) HeaderText(text string) {
	p.header.text = text
}

func (p *presentation) FooterText(text string) {
	p.footer.text = text
}

func (p *presentation) AddScreen(src Screen) {
	p.screens = append(p.screens, src)
}

func (p *presentation) OutputFilePath(path string) {
	p.outputPath = path
}

func (p *presentation) getScreenRenderData() []*renderData {
	screenRenderData := []*renderData{}
	for _, s := range p.screens {
		currentRenderData := &renderData{
			headerText: p.header.text,
			headerPng:  p.header.png,
			footerText: p.footer.text,
			headline:   s.(*screen).headline,
			text:       s.(*screen).text,
			png:        s.(*screen).pngBase64}
		screenRenderData = append(screenRenderData, currentRenderData)
	}
	return screenRenderData
}

func (p *presentation) Render() {
	r := new(renderer)
	r.renderData(p.getScreenRenderData())
	r.outputFilePath(p.outputPath)
	r.render()
}

type renderer struct {
	data       []*renderData
	outputPath string
}

func (r *renderer) renderData(data []*renderData) {
	r.data = data
}

func (r *renderer) outputFilePath(path string) {
	r.outputPath = path
}

func (r *renderer) render() {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 16)
	for _, screenData := range r.data {
		pdf.AddPage()
		pdf.Cell(40, 10, screenData.text)
	}
	err := pdf.OutputFileAndClose(r.outputPath)
	if err != nil {
		fmt.Println(err)
	}
}

type renderData struct {
	headerText string
	headerPng  string
	footerText string
	headline   string
	text       string
	png        string
}

type header struct {
	text string
	png  string
}

type footer struct {
	text string
}

func NewPresentation() *presentation {
	p := new(presentation)
	p.header = new(header)
	p.footer = new(footer)
	p.screens = []Screen{}
	return p
}

type screen struct {
	pngBase64 string
	headline  string
	text      string
}

func (s *screen) Headline(headline string) {
	s.headline = headline
}

func (s *screen) Text(text string) {
	s.text = text
}

func (s *screen) PNGbase64(pngBase64 string) {
	s.pngBase64 = pngBase64
}

func NewScreen() *screen {
	return new(screen)
}
