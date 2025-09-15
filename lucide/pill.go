package lucide

import x "github.com/plainkit/blox"

// Pill creates a Pill Lucide icon.
func Pill(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pill", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10.5 20.5 10-10a4.95 4.95 0 1 0-7-7l-10 10a4.95 4.95 0 1 0 7 7Z"))),
		x.Child(x.Path(x.D("m8.5 8.5 7 7"))),
	)
	return x.Svg(svgArgs...)
}
