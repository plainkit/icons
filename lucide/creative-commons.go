package lucide

import x "github.com/plainkit/html"

// CreativeCommons creates a Creative Commons Lucide icon.
func CreativeCommons(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-creative-commons", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M10 9.3a2.8 2.8 0 0 0-3.5 1 3.1 3.1 0 0 0 0 3.4 2.7 2.7 0 0 0 3.5 1"))),
		x.Child(x.Path(x.D("M17 9.3a2.8 2.8 0 0 0-3.5 1 3.1 3.1 0 0 0 0 3.4 2.7 2.7 0 0 0 3.5 1"))),
	)
	return x.Svg(svgArgs...)
}
