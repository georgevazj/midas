package main

type Note struct {
	To string `xml:"to" yaml:"to" json:"to"`
	From string `xml:"from" yaml:"from" json:"from"`
	Heading string `xml:"heading" yaml:"heading" json:"heading"`
	Body string `xml:"body" yaml:"body" json:"body"`
}

type Config struct {
	Baseurl   string `yml:"baseUrl"`
	Title     string `yml:"title"`
	Templates string `yml:"templates"`
	Posts     string `yml:"posts"`
	Public    string `yml:"public"`
	Admin     string `yml:"admin"`
	Metadata  string `yml:"metadata"`
	Index     string `yml:"index"`
}