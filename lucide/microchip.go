package lucide

import x "github.com/bloxui/blox"

// Microchip creates a Microchip Lucide icon.
func Microchip(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-microchip", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 12h2"))),
		x.Child(x.Path(x.D("M18 16h2"))),
		x.Child(x.Path(x.D("M18 20h2"))),
		x.Child(x.Path(x.D("M18 4h2"))),
		x.Child(x.Path(x.D("M18 8h2"))),
		x.Child(x.Path(x.D("M4 12h2"))),
		x.Child(x.Path(x.D("M4 16h2"))),
		x.Child(x.Path(x.D("M4 20h2"))),
		x.Child(x.Path(x.D("M4 4h2"))),
		x.Child(x.Path(x.D("M4 8h2"))),
		x.Child(x.Path(x.D("M8 2a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2h-1.5c-.276 0-.494.227-.562.495a2 2 0 0 1-3.876 0C9.994 2.227 9.776 2 9.5 2z"))),
	)
	return x.Svg(svgArgs...)
}
