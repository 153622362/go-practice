package model

import (
	"go-practice/crawler/engine"
)

type SearchResult struct {
	Hits     int64 //共
	Start    int   //开始位置
	Query    string
	PrevFrom int
	NextFrom int
	Items    []engine.Item
}
