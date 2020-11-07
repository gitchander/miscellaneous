package main

// Background ColorRGBf

type Palette struct {
	Foreground ColorRGBf
	Path       ColorRGBf // Trajectory

	MassFill   ColorRGBf
	MassStroke ColorRGBf
}

var palettes = []Palette{
	{
		Foreground: Black,
		Path:       RGBf(0.3, 0.3, 1),

		MassFill:   Blue,
		MassStroke: Black,
	},
	{
		Foreground: Black,
		Path:       RGBf(0.3, 0.3, 1),

		MassFill:   Red,
		MassStroke: Black,
	},
	{
		Foreground: Black,
		Path:       RGBf(0.3, 0.3, 1),

		MassFill:   Yellow,
		MassStroke: Black,
	},
	{
		Foreground: Black,
		Path:       RGBf(0.3, 0.3, 1),

		MassFill:   Green,
		MassStroke: Black,
	},
}

func GetPalette(i int) Palette {
	return palettes[mod(i, len(palettes))]
}
