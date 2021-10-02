package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[int]int
	FloatMap  map[float64]float64
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
