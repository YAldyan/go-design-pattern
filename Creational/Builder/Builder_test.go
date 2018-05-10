package Builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	/* Director */
	manufacturingComplex := ManufacturingDirector{}

	/*kriteria builder*/
	carBuilder := &CarBuilder{}

	/*Director membentuk Builder*/
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	/*Builder menghasilkan objek*/
	car := carBuilder.Build()

	/*kriteria objek*/
	if car.Wheels != 4 {
		t.Error("Roda untuk mobil haruslah 4, bukan ", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Error("Struktur untuk mobil haruslah 'Car', bukan ", car.Structure)
	}

	if car.Seats != 5 {
		t.Error("Kursi untuk mobil haruslah 5, bukan ", car.Seats)
	}

	/*kriteria objek*/
	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n",
			motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n",
			motorbike.Structure)
	}
}
