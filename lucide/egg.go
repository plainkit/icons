package lucide

import x "github.com/bloxui/blox"

// Egg creates a Egg Lucide icon.
func Egg(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-egg", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2C8 2 4 8 4 14a8 8 0 0 0 16 0c0-6-4-12-8-12"))),
	)
	return x.Svg(svgArgs...)
}
