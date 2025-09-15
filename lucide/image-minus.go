package lucide

import x "github.com/plainkit/blox"

// ImageMinus creates a Image Minus Lucide icon.
func ImageMinus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-image-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7"))),
		x.Child(x.Line(x.X1("16"), x.X2("22"), x.Y1("5"), x.Y2("5"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("9"), x.R("2"))),
		x.Child(x.Path(x.D("m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21"))),
	)
	return x.Svg(svgArgs...)
}
