package main

type Palette struct {
	clBackground ColorRGBf
	clForeground ColorRGBf
	clPath       ColorRGBf // Trajectory

	clMassFill   ColorRGBf
	clMassStroke ColorRGBf
}

func makePalette1() Palette {
	return Palette{
		clBackground: RGBf(1, 1, 1),
		clForeground: RGBf(0, 0, 0),
		clPath:       RGBf(0.3, 0.3, 1),

		clMassFill:   RGBf(0, 0.5, 0),
		clMassStroke: RGBf(0, 0.2, 0),
	}
}
