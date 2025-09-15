package lucide

import x "github.com/plainkit/html"

// CirclePlus creates a Circle Plus Lucide icon.
func CirclePlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M8 12h8"))),
		x.Child(x.Path(x.D("M12 8v8"))),
	)
	return x.Svg(svgArgs...)
}
