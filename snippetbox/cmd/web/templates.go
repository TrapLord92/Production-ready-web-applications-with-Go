package main

import "github.com/TrapLord92/Production-ready-web-applications-with-Go/internal/models"

// Include a Snippets field in the templateData struct.
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
