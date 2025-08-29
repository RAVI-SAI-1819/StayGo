package models

// TemplateData holds dynamic data passed to HTML templates
type TemplateData struct {
	StringMap map[string]string      // Key-value string pairs
	IntMap    map[string]int         // Key-value int pairs
	FloatMap  map[string]float32     // Key-value float pairs
	Data      map[string]interface{} // Generic data
	CSRFToken string                 // CSRF token for forms
	Flash     string                 // Flash message
	Warning   string                 // Warning message
	Error     string                 // Error message
}