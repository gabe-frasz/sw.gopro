package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/gabe-frasz/qslsp/internal/analysis"
	"github.com/gabe-frasz/qslsp/internal/lsp"
	"github.com/gabe-frasz/qslsp/internal/rpc"
)

func main() {
	logger := getLogger("/home/bielsz/Dev/sw.gopro/qƨlsp/log.txt")
	logger.Println("Starting qƨlsp")

	state := analysis.NewState()
	writer := os.Stdout

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Println("Got an error while decoding message:", err)
			continue
		}

		handleMessage(logger, writer, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state *analysis.State, method string, contents []byte) {
	logger.Printf("Received message with method '%s'\n", method)

	switch method {
	case "initialize":
		var req lsp.InitializeRequest
		if err := json.Unmarshal(contents, &req); err != nil {
			logger.Println("Error unmarshalling 'initialize':", err)
			return
		}

		logger.Printf("Connected to %s %s\n", req.Params.ClientInfo.Name, req.Params.ClientInfo.Version)

		res := lsp.NewInitializeResponse(req.ID)
		writeResponse(writer, res)
		logger.Printf("Sent response with id %d\n", res.ID)
	case "textDocument/didOpen":
		var req lsp.TextDocumentDidOpenNotification
		if err := json.Unmarshal(contents, &req); err != nil {
			logger.Println("Error unmarshalling 'textDocument/didOpen':", err)
			return
		}

		diagnostics := state.OpenDocument(req.Params.TextDocument.URI, req.Params.TextDocument.Text)
		logger.Printf("Opened document with uri '%s'\n", req.Params.TextDocument.URI)

		if len(diagnostics) > 0 {
			res := lsp.NewTextDocumentPublishDiagnosticsNotification(req.Params.TextDocument.URI, diagnostics)
			writeResponse(writer, res)
			logger.Printf("Sent diagnostics for document with uri '%s'\n", req.Params.TextDocument.URI)
		}
	case "textDocument/didChange":
		var req lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &req); err != nil {
			logger.Println("Error unmarshalling 'textDocument/didChange':", err)
			return
		}

		for _, change := range req.Params.ContentChanges {
			diagnostics := state.UpdateDocument(req.Params.TextDocument.URI, change.Text)
			logger.Printf("Changed document with uri '%s'\n", req.Params.TextDocument.URI)

			res := lsp.NewTextDocumentPublishDiagnosticsNotification(req.Params.TextDocument.URI, diagnostics)
			writeResponse(writer, res)
			logger.Printf("Sent diagnostics for document with uri '%s'\n", req.Params.TextDocument.URI)
		}

		logger.Printf("Changed document with uri '%s'\n", req.Params.TextDocument.URI)
	case "textDocument/hover":
		var req lsp.TextDocumentHoverRequest
		if err := json.Unmarshal(contents, &req); err != nil {
			logger.Println("Error unmarshalling 'textDocument/hover':", err)
			return
		}

		contents := state.Hover(req.Params.TextDocument.URI, req.Params.Position)
		res := lsp.NewTextDocumentHoverResponse(req.ID, contents)
		writeResponse(writer, res)
		logger.Printf("Sent hover response for document with uri %s\n", req.Params.TextDocument.URI)
	case "textDocument/definition":
		var req lsp.TextDocumentDefinitionRequest
		if err := json.Unmarshal(contents, &req); err != nil {
			logger.Println("Error unmarshalling 'textDocument/definition':", err)
			return
		}

		location := state.GoToDefinition(req.Params.TextDocument.URI, req.Params.Position)
		res := lsp.NewTextDocumentDefinitionResponse(req.ID, location)
		writeResponse(writer, res)
		logger.Printf("Sent definition for document with uri %s\n", req.Params.TextDocument.URI)
	case "textDocument/codeAction":
		var req lsp.TextDocumentCodeActionRequest
		if err := json.Unmarshal(contents, &req); err != nil {
			logger.Println("Error unmarshalling 'textDocument/codeAction':", err)
			return
		}

		actions := state.GetCodeActions(req.Params.TextDocument.URI)
		res := lsp.NewTextDocumentCodeActionResponse(req.ID, actions)
		writeResponse(writer, res)
		logger.Printf("Sent code actions for document with uri %s\n", req.Params.TextDocument.URI)
	case "textDocument/completion":
		var req lsp.TextDocumentCompletionRequest
		if err := json.Unmarshal(contents, &req); err != nil {
			logger.Println("Error unmarshalling 'textDocument/completion':", err)
			return
		}

		items := state.GetCompletions(req.Params.TextDocumentPositionParams.TextDocument.URI)
		res := lsp.NewTextDocumentCompletionResponse(req.ID, items)
		writeResponse(writer, res)
		logger.Printf("Sent completions for document with uri %s\n", req.Params.TextDocument.URI)
	case "shutdown":
		logger.Println("Shutting down")
		os.Exit(0)
	}
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logFile, "[qƨlsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func writeResponse(writer io.Writer, msg any) {
	res := rpc.EncodeMessage(msg)
	writer.Write([]byte(res))
}
