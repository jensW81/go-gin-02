package go_gin

type (
	Kennzeichen string
	Auto        interface {
		MotorStarten() error
		Beschleunigen()
		Bremsen()
	}

	Audi struct {
		Kennzeichen          Kennzeichen `json:"kennzeichen"`
		Getriebe             *Getriebe   `json:"getriebe,omitempty"`
		GeschwindigkeitInKmH int         `json:"-"`
		motorGestartet       bool        `json:"-"`
	}

	Getriebe struct {
		EingangsDrehzahl int
		EingelegterGang  int
	}
)

func (a Audi) MotorStarten() error {
	//TODO implement me
	panic("implement me")
}

func (a Audi) Beschleunigen() {
	//TODO implement me
	panic("implement me")
}

func (a Audi) Bremsen() {
	//TODO implement me
	panic("implement me")
}
