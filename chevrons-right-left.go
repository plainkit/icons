package lucide

import x "github.com/bloxui/blox"

// ChevronsRightLeft creates a Chevrons Right Left Lucide icon.
func ChevronsRightLeft(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chevrons-right-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m20 17-5-5 5-5"))),
		x.Child(x.Path(x.D("m4 17 5-5-5-5"))),
	)
	return x.Svg(svgArgs...)
}
