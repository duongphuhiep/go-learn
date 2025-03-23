package main

import "time"

type Entry[T any] struct {
	Value *T
}

type ComplexType struct {
	ID         uint64
	Name       string
	Age        int32
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Addresses  []Address
	Phone      string
	Email      string
	Website    string
	Latitude   float64
	Longitude  float64
	Tags       []string
	Categories []Category
	Score      float32
	Rating     int16
	Reviews    []Review
	Images     []Image
	Videos     []Video
	Metadata   map[string]string
}

type Address struct {
	Street  string
	City    string
	State   string
	Zip     string
	Country string
}

type Category struct {
	ID   uint64
	Name string
}

type Review struct {
	ID      uint64
	Rating  int16
	Comment string
}

type Image struct {
	URL    string
	Width  int32
	Height int32
}

type Video struct {
	URL  string
	Type string
}
