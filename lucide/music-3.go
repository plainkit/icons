package lucide

import x "github.com/plainkit/html"

// Music3 creates a Music 3 Lucide icon.
func Music3(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-music-3", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("18"), x.R("4"))),
		x.Child(x.Path(x.D("M16 18V2"))),
	)
	return x.Svg(svgArgs...)
}
