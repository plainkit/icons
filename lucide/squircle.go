package lucide

import x "github.com/bloxui/blox"

// Squircle creates a Squircle Lucide icon.
func Squircle(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-squircle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3c7.2 0 9 1.8 9 9s-1.8 9-9 9-9-1.8-9-9 1.8-9 9-9"))),
	)
	return x.Svg(svgArgs...)
}
