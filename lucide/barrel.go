package lucide

import x "github.com/bloxui/blox"

// Barrel creates a Barrel Lucide icon.
func Barrel(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-barrel", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 3a41 41 0 0 0 0 18"))),
		x.Child(x.Path(x.D("M14 3a41 41 0 0 1 0 18"))),
		x.Child(x.Path(x.D("M17 3a2 2 0 0 1 1.68.92 15.25 15.25 0 0 1 0 16.16A2 2 0 0 1 17 21H7a2 2 0 0 1-1.68-.92 15.25 15.25 0 0 1 0-16.16A2 2 0 0 1 7 3z"))),
		x.Child(x.Path(x.D("M3.84 17h16.32"))),
		x.Child(x.Path(x.D("M3.84 7h16.32"))),
	)
	return x.Svg(svgArgs...)
}
