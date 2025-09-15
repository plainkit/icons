package lucide

import x "github.com/plainkit/html"

// Flashlight creates a Flashlight Lucide icon.
func Flashlight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flashlight", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 6c0 2-2 2-2 4v10a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V10c0-2-2-2-2-4V2h12z"))),
		x.Child(x.Line(x.X1("6"), x.X2("18"), x.Y1("6"), x.Y2("6"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("12"), x.Y2("12"))),
	)
	return x.Svg(svgArgs...)
}
