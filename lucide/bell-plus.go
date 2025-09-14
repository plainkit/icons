package lucide

import x "github.com/bloxui/blox"

// BellPlus creates a Bell Plus Lucide icon.
func BellPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bell-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.268 21a2 2 0 0 0 3.464 0"))),
		x.Child(x.Path(x.D("M15 8h6"))),
		x.Child(x.Path(x.D("M18 5v6"))),
		x.Child(x.Path(x.D("M20.002 14.464a9 9 0 0 0 .738.863A1 1 0 0 1 20 17H4a1 1 0 0 1-.74-1.673C4.59 13.956 6 12.499 6 8a6 6 0 0 1 8.75-5.332"))),
	)
	return x.Svg(svgArgs...)
}
