package lucide

import x "github.com/bloxui/blox"

// Bolt creates a Bolt Lucide icon.
func Bolt(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-bolt", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("4"))),
	)
	return x.Svg(svgArgs...)
}
