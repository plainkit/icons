package lucide

import x "github.com/bloxui/blox"

// Vote creates a Vote Lucide icon.
func Vote(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-vote", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m9 12 2 2 4-4"))),
		x.Child(x.Path(x.D("M5 7c0-1.1.9-2 2-2h10a2 2 0 0 1 2 2v12H5V7Z"))),
		x.Child(x.Path(x.D("M22 19H2"))),
	)
	return x.Svg(svgArgs...)
}
