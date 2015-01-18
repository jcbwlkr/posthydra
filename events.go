package posthydra

type Event struct {
	Title    string
	Body     string
	Location string
	Start    string
	End      string
	URL      string
}

type Reader interface {
	Read() ([]*Event, error)
}
