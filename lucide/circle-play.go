package lucide

import x "github.com/bloxui/blox"

// CirclePlay creates a Circle Play Lucide icon.
func CirclePlay(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-play", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 9.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997A1 1 0 0 1 9 14.996z"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
