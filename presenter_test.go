package presenter

import (
	"path"
	"runtime"
	"testing"

	"github.com/ingmardrewing/fs"
)

func tearDown() {
	dir := path.Join(getTestFileDirPath(), "testResources/createdPdfs")
	fs.RemoveDirContents(dir)
}

func getTestFileDirPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func TestNewPresentation(t *testing.T) {
	p := NewPresentation()

	headerText := "test header"
	p.HeaderText(headerText)

	headerPNG := "a base64 encoded png"
	p.HeaderPNGbase64(headerPNG)

	footerText := "test footer"
	p.FooterText(footerText)

	outputPath := "test footer"
	p.OutputFilePath(outputPath)

	if p.header.text != headerText {
		t.Error("Expected", p.header.text, "to be", headerText)
	}

	if p.header.png != headerPNG {
		t.Error("Expected", p.header.text, "to be", headerText)
	}

	if p.header.png != headerPNG {
		t.Error("Expected", p.header.text, "to be", headerText)
	}

	if p.outputPath != outputPath {
		t.Error("Expected", p.outputPath, "to be", outputPath)
	}
}

func TestNewScreen(t *testing.T) {
	s := NewScreen()

	headline := "screen headline"
	s.Headline(headline)

	text := "screen text"
	s.Text(text)

	pngBase64 := "png base 64"
	s.PNGbase64(pngBase64)

	if s.headline != headline {
		t.Error("Expected", s.headline, "to be", headline)
	}

	if s.text != text {
		t.Error("Expected", s.text, "to be", text)
	}

	if s.pngBase64 != pngBase64 {
		t.Error("Expected", s.pngBase64, "to be", pngBase64)
	}
}

func TestGetScreenRenderData(t *testing.T) {
	p := NewPresentation()

	headerText := "test header"
	p.HeaderText(headerText)

	headerPng := "a base64 encoded png"
	p.HeaderPNGbase64(headerPng)

	footerText := "test footer"
	p.FooterText(footerText)

	outputPath := "test footer"
	p.OutputFilePath(outputPath)

	s1 := NewScreen()

	s1Headline := "s1 Headline"
	s1.Headline(s1Headline)

	s1Text := "s1 Text"
	s1.Text(s1Text)

	s1Png := "s1 Png"
	s1.PNGbase64(s1Png)

	s2 := NewScreen()

	s2Headline := "s2 Headline"
	s2.Headline(s2Headline)

	s2Text := "s2 Text"
	s2.Text(s2Text)

	s2Png := "s2 Png"
	s2.PNGbase64(s2Png)

	p.AddScreen(s1)
	p.AddScreen(s2)

	screenRenderData := p.getScreenRenderData()

	expectedAmount := 2
	if len(screenRenderData) != expectedAmount {
		t.Error("Excepted to find", expectedAmount, "screens, but didn't.")
	}

	if screenRenderData[0].headerPng != headerPng {
		t.Error("Expected", screenRenderData[0].headerPng, "to be", headerPng)
	}

	if screenRenderData[1].png != s2Png {
		t.Error("Expected", screenRenderData[1].png, "to be", s2Png)
	}

	if screenRenderData[0].text != s1Text {
		t.Error("Expected", screenRenderData[0].text, "to be", s2Text)
	}
}

func TestPresentationRender(t *testing.T) {
	p := NewPresentation()

	headerText := "test header"
	p.HeaderText(headerText)

	headerPng := "a base64 encoded png"
	p.HeaderPNGbase64(headerPng)

	footerText := "test footer"
	p.FooterText(footerText)

	outputPath := "./testResources/createdPdfs/test.pdf"
	p.OutputFilePath(outputPath)

	s1 := NewScreen()

	s1Headline := "s1 Headline"
	s1.Headline(s1Headline)

	s1Text := "s1 Text"
	s1.Text(s1Text)

	s1Png := "s1 Png"
	s1.PNGbase64(s1Png)

	s2 := NewScreen()

	s2Headline := "s2 Headline"
	s2.Headline(s2Headline)

	s2Text := "s2 Text"
	s2.Text(s2Text)

	s2Png := "s2 Png"
	s2.PNGbase64(s2Png)

	p.AddScreen(s1)
	p.AddScreen(s2)
	p.Render()

	pdfPath := path.Join(getTestFileDirPath(), outputPath)
	pdfFileExists, _ := fs.PathExists(pdfPath)

	if !pdfFileExists {
		t.Error("Expected to find generated pdf at", pdfPath)
	}
}
