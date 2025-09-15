package lucide

import x "github.com/plainkit/blox"

// Flame creates a Flame Lucide icon.
func Flame(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flame", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3q1 4 4 6.5t3 5.5a1 1 0 0 1-14 0 5 5 0 0 1 1-3 1 1 0 0 0 5 0c0-2-1.5-3-1.5-5q0-2 2.5-4"))),
	)
	return x.Svg(svgArgs...)
}
