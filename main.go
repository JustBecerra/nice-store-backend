package main

type product struct {
    id     int  `json:"id"`
    title  string  `json:"title"`
    price  float64 `json:"price"`
	description string `json:"description"`
	category string `json:"category"`
}