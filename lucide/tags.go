package lucide

import x "github.com/plainkit/html"

// Tags creates a Tags Lucide icon.
func Tags(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tags", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13.172 2a2 2 0 0 1 1.414.586l6.71 6.71a2.4 2.4 0 0 1 0 3.408l-4.592 4.592a2.4 2.4 0 0 1-3.408 0l-6.71-6.71A2 2 0 0 1 6 9.172V3a1 1 0 0 1 1-1z"))),
		x.Child(x.Path(x.D("M2 7v6.172a2 2 0 0 0 .586 1.414l6.71 6.71a2.4 2.4 0 0 0 3.191.193"))),
		x.Child(x.Circle(x.Cx("10.5"), x.Cy("6.5"), x.R(".5"), x.Fill("currentColor"))),
	)
	return x.Svg(svgArgs...)
}
