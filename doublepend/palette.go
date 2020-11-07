package main

type Palette struct {
	Background ColorRGBf
	Foreground ColorRGBf
	Path       ColorRGBf // Trajectory

	MassFill   ColorRGBf
	MassStroke ColorRGBf
}

var palettes = []Palette{
	{
		Background: RGBf(1, 1, 1),
		Foreground: RGBf(0, 0, 0),
		Path:       RGBf(0.3, 0.3, 1),

		MassFill:   RGBf(0.8, 0.8, 0),
		MassStroke: RGBf(0.2, 0.2, 0),
	},
	{
		Background: RGBf(1, 1, 1),
		Foreground: RGBf(0, 0, 0),
		Path:       RGBf(0.3, 0.3, 1),

		MassFill:   RGBf(0, 0.5, 0),
		MassStroke: RGBf(0, 0.2, 0),
	},
}
