package lucide

import x "github.com/plainkit/blox"

// Redo2 creates a Redo 2 Lucide icon.
func Redo2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-redo-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 14 5-5-5-5"))),
		x.Child(x.Path(x.D("M20 9H9.5A5.5 5.5 0 0 0 4 14.5A5.5 5.5 0 0 0 9.5 20H13"))),
	)
	return x.Svg(svgArgs...)
}
