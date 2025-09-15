package lucide

import x "github.com/plainkit/html"

// Undo2 creates a Undo 2 Lucide icon.
func Undo2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-undo-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 14 4 9l5-5"))),
		x.Child(x.Path(x.D("M4 9h10.5a5.5 5.5 0 0 1 5.5 5.5a5.5 5.5 0 0 1-5.5 5.5H11"))),
	)
	return x.Svg(svgArgs...)
}
