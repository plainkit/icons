package lucide

import x "github.com/plainkit/blox"

// Unlink2 creates a Unlink 2 Lucide icon.
func Unlink2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-unlink-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 7h2a5 5 0 0 1 0 10h-2m-6 0H7A5 5 0 0 1 7 7h2"))),
	)
	return x.Svg(svgArgs...)
}
