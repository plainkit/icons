package lucide

import x "github.com/plainkit/html"

// TextSearch creates a Text Search Lucide icon.
func TextSearch(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-text-search", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 5H3"))),
		x.Child(x.Path(x.D("M10 12H3"))),
		x.Child(x.Path(x.D("M10 19H3"))),
		x.Child(x.Circle(x.Cx("17"), x.Cy("15"), x.R("3"))),
		x.Child(x.Path(x.D("m21 19-1.9-1.9"))),
	)
	return x.Svg(svgArgs...)
}
