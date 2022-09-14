package calculators

type IAngelMonth struct {
	Numberic     int
	Alphabetical string
}

type IAngelAnswer struct {
	PreAngel int
	Angel    int
}

type IAngel struct {
	At       string
	FullDate string
	Day      int
	Month    *IAngelMonth
	Year     int
	Answer   *IAngelAnswer
}

func Angel() *IAngel {

}
