package main

type Input struct {
	Name     string           `json:"name"`
	Contract string           `json:"contract"`
	Claims   map[string]Proof `json:"claims"`
}

type Proof struct {
	Index  uint64   `json:"index"`
	Amount string   `json:"amount"`
	Proof  []string `json:"proof"`
}

type Contract struct {
	Name     string `json:"name"`
	Contract string `json:"contract"`
}
