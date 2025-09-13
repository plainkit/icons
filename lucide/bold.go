package lucide

import x "github.com/bloxui/blox"

// Bold creates a Bold Lucide icon.
func Bold(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-bold", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 12h9a4 4 0 0 1 0 8H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h7a4 4 0 0 1 0 8"))),
	)
	return x.Svg(svgArgs...)
}
