package lucide

import x "github.com/bloxui/blox"

// Vegan creates a Vegan Lucide icon.
func Vegan(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-vegan", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 8q6 0 6-6-6 0-6 6"))),
		x.Child(x.Path(x.D("M17.41 3.59a10 10 0 1 0 3 3"))),
		x.Child(x.Path(x.D("M2 2a26.6 26.6 0 0 1 10 20c.9-6.82 1.5-9.5 4-14"))),
	)
	return x.Svg(svgArgs...)
}
