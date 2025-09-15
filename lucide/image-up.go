package lucide

import x "github.com/plainkit/html"

// ImageUp creates a Image Up Lucide icon.
func ImageUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-image-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.3 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10l-3.1-3.1a2 2 0 0 0-2.814.014L6 21"))),
		x.Child(x.Path(x.D("m14 19.5 3-3 3 3"))),
		x.Child(x.Path(x.D("M17 22v-5.5"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("9"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
