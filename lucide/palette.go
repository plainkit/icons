package lucide

import x "github.com/bloxui/blox"

// Palette creates a Palette Lucide icon.
func Palette(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-palette", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 22a1 1 0 0 1 0-20 10 9 0 0 1 10 9 5 5 0 0 1-5 5h-2.25a1.75 1.75 0 0 0-1.4 2.8l.3.4a1.75 1.75 0 0 1-1.4 2.8z"))),
		x.Child(x.Circle(x.Cx("13.5"), x.Cy("6.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Circle(x.Cx("17.5"), x.Cy("10.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Circle(x.Cx("6.5"), x.Cy("12.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Circle(x.Cx("8.5"), x.Cy("7.5"), x.R(".5"), x.Fill("currentColor"))),
	)
	return x.Svg(svgArgs...)
}
