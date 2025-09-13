package lucide

import x "github.com/bloxui/blox"

// Turtle creates a Turtle Lucide icon.
func Turtle(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-turtle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m12 10 2 4v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3a8 8 0 1 0-16 0v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3l2-4h4Z"))),
		x.Child(x.Path(x.D("M4.82 7.9 8 10"))),
		x.Child(x.Path(x.D("M15.18 7.9 12 10"))),
		x.Child(x.Path(x.D("M16.93 10H20a2 2 0 0 1 0 4H2"))),
	)
	return x.Svg(svgArgs...)
}
