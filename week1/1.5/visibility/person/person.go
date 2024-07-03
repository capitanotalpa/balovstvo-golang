package person

var (
	Public  = 1
	private = 2
)

type Person struct {
	ID     int
	Name   string
	secret string
}

func (p Person) UpdateSecret(secret string) {
	p.secret = secret
}
