package lucide

import x "github.com/bloxui/blox"

// Utensils creates a Utensils Lucide icon.
func Utensils(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-utensils", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 2v7c0 1.1.9 2 2 2h4a2 2 0 0 0 2-2V2"))),
		x.Child(x.Path(x.D("M7 2v20"))),
		x.Child(x.Path(x.D("M21 15V2a5 5 0 0 0-5 5v6c0 1.1.9 2 2 2h3Zm0 0v7"))),
	)
	return x.Svg(svgArgs...)
}
