package lucide

import x "github.com/plainkit/html"

// Bed creates a Bed Lucide icon.
func Bed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 4v16"))),
		x.Child(x.Path(x.D("M2 8h18a2 2 0 0 1 2 2v10"))),
		x.Child(x.Path(x.D("M2 17h20"))),
		x.Child(x.Path(x.D("M6 8v9"))),
	)
	return x.Svg(svgArgs...)
}
