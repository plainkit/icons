package lucide

import x "github.com/bloxui/blox"

// Infinity creates a Infinity Lucide icon.
func Infinity(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-infinity", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 16c5 0 7-8 12-8a4 4 0 0 1 0 8c-5 0-7-8-12-8a4 4 0 1 0 0 8"))),
	)
	return x.Svg(svgArgs...)
}
