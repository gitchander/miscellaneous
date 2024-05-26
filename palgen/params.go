package palgen

type Params struct {
	A, B, C, D Vec3
}

var ps = []Params{
	Params{
		A: Vec3{0.5, 0.5, 0.5},
		B: Vec3{0.5, 0.5, 0.5},
		C: Vec3{1.0, 1.0, 1.0},
		D: Vec3{0.00, 0.33, 0.67},
	},
	Params{
		A: Vec3{0.5, 0.5, 0.5},
		B: Vec3{0.5, 0.5, 0.5},
		C: Vec3{1.0, 1.0, 1.0},
		D: Vec3{0.00, 0.10, 0.20},
	},
	Params{
		A: Vec3{0.5, 0.5, 0.5},
		B: Vec3{0.5, 0.5, 0.5},
		C: Vec3{1.0, 1.0, 1.0},
		D: Vec3{0.30, 0.20, 0.20},
	},
	Params{
		A: Vec3{0.5, 0.5, 0.5},
		B: Vec3{0.5, 0.5, 0.5},
		C: Vec3{1.0, 1.0, 0.5},
		D: Vec3{0.80, 0.90, 0.30},
	},
	Params{
		A: Vec3{0.5, 0.5, 0.5},
		B: Vec3{0.5, 0.5, 0.5},
		C: Vec3{1.0, 0.7, 0.4},
		D: Vec3{0.00, 0.15, 0.20},
	},
}
