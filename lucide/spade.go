package lucide

import x "github.com/plainkit/html"

// Spade creates a Spade Lucide icon.
func Spade(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-spade", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 18v4"))),
		x.Child(x.Path(x.D("M2 14.499a5.5 5.5 0 0 0 9.591 3.675.6.6 0 0 1 .818.001A5.5 5.5 0 0 0 22 14.5c0-2.29-1.5-4-3-5.5l-5.492-5.312a2 2 0 0 0-3-.02L5 8.999c-1.5 1.5-3 3.2-3 5.5"))),
	)
	return x.Svg(svgArgs...)
}
