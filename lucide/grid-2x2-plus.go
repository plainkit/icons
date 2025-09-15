package lucide

import x "github.com/plainkit/html"

// Grid2x2Plus creates a Grid 2x2 Plus Lucide icon.
func Grid2x2Plus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-grid-2x2-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3v17a1 1 0 0 1-1 1H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v6a1 1 0 0 1-1 1H3"))),
		x.Child(x.Path(x.D("M16 19h6"))),
		x.Child(x.Path(x.D("M19 22v-6"))),
	)
	return x.Svg(svgArgs...)
}
