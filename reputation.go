package main

type Reputation struct {
	SHA256    string `json:"sha256"`
	Rep       string `json:"rep"`
	DateAdded string `json:"date"`
	TTL       int    `json:"ttl"`
}

//type Reputations []Reputation
