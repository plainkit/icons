package lucide

import x "github.com/plainkit/html"

// CircleSlash2 creates a Circle Slash 2 Lucide icon.
func CircleSlash2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-slash-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 2 2 22"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
