package lucide

import x "github.com/bloxui/blox"

// LogOut creates a Log Out Lucide icon.
func LogOut(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-log-out", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 17 5-5-5-5"))),
		x.Child(x.Path(x.D("M21 12H9"))),
		x.Child(x.Path(x.D("M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"))),
	)
	return x.Svg(svgArgs...)
}
