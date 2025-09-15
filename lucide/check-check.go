package lucide

import x "github.com/plainkit/blox"

// CheckCheck creates a Check Check Lucide icon.
func CheckCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-check-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 6 7 17l-5-5"))),
		x.Child(x.Path(x.D("m22 10-7.5 7.5L13 16"))),
	)
	return x.Svg(svgArgs...)
}
