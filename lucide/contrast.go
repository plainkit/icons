package lucide

import x "github.com/plainkit/html"

// Contrast creates a Contrast Lucide icon.
func Contrast(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-contrast", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M12 18a6 6 0 0 0 0-12v12z"))),
	)
	return x.Svg(svgArgs...)
}
