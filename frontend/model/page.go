package model

type SearchResult struct {
	Hits     int64
	//Hits     int
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []interface{}
	//Items    []engine.Item
}
