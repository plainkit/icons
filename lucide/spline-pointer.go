package lucide

import x "github.com/plainkit/html"

// SplinePointer creates a Spline Pointer Lucide icon.
func SplinePointer(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-spline-pointer", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.034 12.681a.498.498 0 0 1 .647-.647l9 3.5a.5.5 0 0 1-.033.943l-3.444 1.068a1 1 0 0 0-.66.66l-1.067 3.443a.5.5 0 0 1-.943.033z"))),
		x.Child(x.Path(x.D("M5 17A12 12 0 0 1 17 5"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("5"), x.R("2"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("19"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
