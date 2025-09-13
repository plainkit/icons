package lucide

import x "github.com/bloxui/blox"

// Rainbow creates a Rainbow Lucide icon.
func Rainbow(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-rainbow", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 17a10 10 0 0 0-20 0"))),
		x.Child(x.Path(x.D("M6 17a6 6 0 0 1 12 0"))),
		x.Child(x.Path(x.D("M10 17a2 2 0 0 1 4 0"))),
	)
	return x.Svg(svgArgs...)
}
