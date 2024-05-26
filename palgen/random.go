package palgen

import (
	"fmt"
	"math/rand"
)

func RandParams(r *rand.Rand, p *Params, n int) {

	p.A = Vec3{0.5, 0.5, 0.5}
	p.B = Vec3{0.5, 0.5, 0.5}

	switch 0 {

	case 0:
		p.C = Vec3{
			0: float64(r.Intn(n)),
			1: float64(r.Intn(n)),
			2: float64(r.Intn(n)),
		}

	case 1:
		p.C = Vec3{
			0: float64(r.Intn(2*n)) / 2,
			1: float64(r.Intn(2*n)) / 2,
			2: float64(r.Intn(2*n)) / 2,
		}
		fmt.Println(p.C)

	case 2:
		p.C = Vec3{
			0: r.Float64() * float64(n),
			1: r.Float64() * float64(n),
			2: r.Float64() * float64(n),
		}
	}

	// Phases
	// random value in range [0..1]
	p.D = Vec3{
		0: r.Float64(),
		1: r.Float64(),
		2: r.Float64(),
	}
}
