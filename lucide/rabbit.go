package lucide

import x "github.com/bloxui/blox"

// Rabbit creates a Rabbit Lucide icon.
func Rabbit(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rabbit", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 16a3 3 0 0 1 2.24 5"))),
		x.Child(x.Path(x.D("M18 12h.01"))),
		x.Child(x.Path(x.D("M18 21h-8a4 4 0 0 1-4-4 7 7 0 0 1 7-7h.2L9.6 6.4a1 1 0 1 1 2.8-2.8L15.8 7h.2c3.3 0 6 2.7 6 6v1a2 2 0 0 1-2 2h-1a3 3 0 0 0-3 3"))),
		x.Child(x.Path(x.D("M20 8.54V4a2 2 0 1 0-4 0v3"))),
		x.Child(x.Path(x.D("M7.612 12.524a3 3 0 1 0-1.6 4.3"))),
	)
	return x.Svg(svgArgs...)
}
