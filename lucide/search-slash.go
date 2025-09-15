package lucide

import x "github.com/plainkit/html"

// SearchSlash creates a Search Slash Lucide icon.
func SearchSlash(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-search-slash", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m13.5 8.5-5 5"))),
		x.Child(x.Circle(x.Cx("11"), x.Cy("11"), x.R("8"))),
		x.Child(x.Path(x.D("m21 21-4.3-4.3"))),
	)
	return x.Svg(svgArgs...)
}
