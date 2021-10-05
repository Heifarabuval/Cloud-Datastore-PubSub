package models

type CustomMap struct {
	Key   string `datastore:"key" json:"key"`
	Value int64  `datastore:"value" json:"value"`
}
