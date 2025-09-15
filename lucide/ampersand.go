package lucide

import x "github.com/plainkit/blox"

// Ampersand creates a Ampersand Lucide icon.
func Ampersand(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ampersand", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17.5 12c0 4.4-3.6 8-8 8A4.5 4.5 0 0 1 5 15.5c0-6 8-4 8-8.5a3 3 0 1 0-6 0c0 3 2.5 8.5 12 13"))),
		x.Child(x.Path(x.D("M16 12h3"))),
	)
	return x.Svg(svgArgs...)
}
