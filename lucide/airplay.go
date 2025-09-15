package lucide

import x "github.com/plainkit/blox"

// Airplay creates a Airplay Lucide icon.
func Airplay(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-airplay", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 17H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1"))),
		x.Child(x.Path(x.D("m12 15 5 6H7Z"))),
	)
	return x.Svg(svgArgs...)
}
