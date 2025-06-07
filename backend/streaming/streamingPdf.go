package streaming

import (
	"app-profesor/backend/server"
	"encoding/json"

	"github.com/google/uuid"
)

type PDFStream struct {
	pageClients map[string]func(int)
	pdfClients  map[string]func([]byte)
	currentPDF  []byte
}

func NewPDFStream() *PDFStream {
	return &PDFStream{
		pageClients: make(map[string]func(int)),
	}
}

func (p *PDFStream) BroadcastPage(page int) {
	msg := map[string]interface{}{
		"type": "pdf_page",
		"page": page,
	}
	data, _ := json.Marshal(msg)
	server.SendToAll(data)
}

func (p *PDFStream) OnPageUpdate(cb func(int)) {
	id := uuid.New().String()
	p.pageClients[id] = cb
}

func (p *PDFStream) BroadcastPDF(content []byte) {
	msg := map[string]interface{}{
		"type": "pdf_init",
		"data": content,
	}
	data, _ := json.Marshal(msg)
	server.SendToAll(data)
}

func (p *PDFStream) OnPDFUpdate(cb func([]byte)) {
	id := uuid.New().String()
	p.pdfClients[id] = cb
	cb(p.currentPDF) // Enviar el PDF actual al suscribirse
}
