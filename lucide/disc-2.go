package lucide

import x "github.com/plainkit/html"

// Disc2 creates a Disc 2 Lucide icon.
func Disc2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-disc-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("4"))),
		x.Child(x.Path(x.D("M12 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
