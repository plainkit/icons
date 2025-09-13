package lucide

import x "github.com/bloxui/blox"

// LogIn creates a Log In Lucide icon.
func LogIn(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-log-in", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10 17 5-5-5-5"))),
		x.Child(x.Path(x.D("M15 12H3"))),
		x.Child(x.Path(x.D("M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4"))),
	)
	return x.Svg(svgArgs...)
}
