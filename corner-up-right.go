package lucide

import x "github.com/bloxui/blox"

// CornerUpRight creates a Corner Up Right Lucide icon.
func CornerUpRight(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-corner-up-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 14 5-5-5-5"))),
		x.Child(x.Path(x.D("M4 20v-7a4 4 0 0 1 4-4h12"))),
	)
	return x.Svg(svgArgs...)
}
