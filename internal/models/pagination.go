package models

type FiltAndPagin struct {
	FilterMap map[string]string
	Key       string
	Values    []interface{}
	Where     []string
	Limit     int
	Offset    int
}
