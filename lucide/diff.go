package lucide

import x "github.com/plainkit/html"

// Diff creates a Diff Lucide icon.
func Diff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-diff", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3v14"))),
		x.Child(x.Path(x.D("M5 10h14"))),
		x.Child(x.Path(x.D("M5 21h14"))),
	)
	return x.Svg(svgArgs...)
}
