package ex12

type XKCD struct {
	Number     int    `json:"num"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	URL        string `json:"img"`
}

type Index struct {
	Entries []XKCD `json:"entries"`
}
