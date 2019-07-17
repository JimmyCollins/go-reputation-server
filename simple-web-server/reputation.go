package main
 
type Reputation struct {
    Value      string	`json:"value"`
    TTL int             `json:"ttl"`
}
 
type Reputations []Reputation