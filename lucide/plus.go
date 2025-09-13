package lucide

import x "github.com/bloxui/blox"

// Plus creates a Plus Lucide icon.
func Plus(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 12h14"))),
		x.Child(x.Path(x.D("M12 5v14"))),
	)
	return x.Svg(svgArgs...)
}
