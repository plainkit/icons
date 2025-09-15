package lucide

import x "github.com/plainkit/html"

// BedDouble creates a Bed Double Lucide icon.
func BedDouble(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bed-double", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 20v-8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8"))),
		x.Child(x.Path(x.D("M4 10V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4"))),
		x.Child(x.Path(x.D("M12 4v6"))),
		x.Child(x.Path(x.D("M2 18h20"))),
	)
	return x.Svg(svgArgs...)
}
