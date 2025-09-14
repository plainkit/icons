package lucide

import x "github.com/bloxui/blox"

// Slash creates a Slash Lucide icon.
func Slash(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-slash", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 2 2 22"))),
	)
	return x.Svg(svgArgs...)
}
