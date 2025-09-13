package lucide

import x "github.com/bloxui/blox"

// ChevronsLeftRight creates a Chevrons Left Right Lucide icon.
func ChevronsLeftRight(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chevrons-left-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m9 7-5 5 5 5"))),
		x.Child(x.Path(x.D("m15 7 5 5-5 5"))),
	)
	return x.Svg(svgArgs...)
}
