package lucide

import x "github.com/bloxui/blox"

// Mountain creates a Mountain Lucide icon.
func Mountain(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mountain", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m8 3 4 8 5-5 5 15H2L8 3z"))),
	)
	return x.Svg(svgArgs...)
}
