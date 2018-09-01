package ex12

type XKCD struct {
	Number     int    `json:"num"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
}

type Index struct {
	Entries []XKCD `json:"entries"`
}
