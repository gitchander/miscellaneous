package main

// Background ColorRGBf
// Path ColorRGBf Trajectory

type Palette struct {
	Foreground ColorRGBf
	MassFill   ColorRGBf
	MassStroke ColorRGBf
}

// var palettes = []Palette{
// 	{
// 		Foreground: Black,
// 		Path:       RGBf(0.3, 0.3, 1),

// 		MassFill:   Blue,
// 		MassStroke: Black,
// 	},
// 	{
// 		Foreground: Black,
// 		Path:       RGBf(0.3, 0.3, 1),

// 		MassFill:   Red,
// 		MassStroke: Black,
// 	},
// 	{
// 		Foreground: Black,
// 		Path:       RGBf(0.3, 0.3, 1),

// 		MassFill:   Yellow,
// 		MassStroke: Black,
// 	},
// 	{
// 		Foreground: Black,
// 		Path:       RGBf(0.3, 0.3, 1),

// 		MassFill:   Green,
// 		MassStroke: Black,
// 	},
// }

var palettes = []Palette{
	{
		Foreground: Gray50,
		MassFill:   hexToRGBf(0xf30f0e), // red
		MassStroke: Gray25,
	},
	{
		Foreground: Gray50,
		MassFill:   hexToRGBf(0x297eb4), // blue
		MassStroke: Gray25,
	},
	{
		Foreground: Gray50,
		MassFill:   hexToRGBf(0xff6f01), // orange
		MassStroke: Gray25,
	},
	{
		Foreground: Gray50,
		MassFill:   hexToRGBf(0x46aa32), // green
		MassStroke: Gray25,
	},
	{
		Foreground: Gray50,
		MassFill:   hexToRGBf(0xcd0525),
		MassStroke: Gray25,
	},
}

func GetPalette(i int) Palette {
	return palettes[mod(i, len(palettes))]
}
