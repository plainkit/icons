package lucide

import x "github.com/plainkit/html"

// View creates a View Lucide icon.
func View(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-view", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 17v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2"))),
		x.Child(x.Path(x.D("M21 7V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("1"))),
		x.Child(x.Path(x.D("M18.944 12.33a1 1 0 0 0 0-.66 7.5 7.5 0 0 0-13.888 0 1 1 0 0 0 0 .66 7.5 7.5 0 0 0 13.888 0"))),
	)
	return x.Svg(svgArgs...)
}
