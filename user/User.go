package user

type age int
type User struct {
	Email  string `json:"email"`
	Age    age    `json:"age"`
	Active bool   `json:"act"`
}

type fio struct {
	f string `json:"f"`
	i string
	o string
}

type geo struct {
	Lat *int
	Lon int
}

func NewGeo(lat int, lon int) *geo {
	return &geo{Lat: &lat, Lon: lon}
}

func NewFio(f string, i string, o string) *fio {

	return &fio{f: f, i: i, o: o}
}

func (name *fio) FirstName() string {
	return name.f
}

func (name *fio) SetFistName(fistName string) {
	name.f = fistName
	//errors.New("essfdfsd")
}
