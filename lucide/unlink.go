package lucide

import x "github.com/plainkit/blox"

// Unlink creates a Unlink Lucide icon.
func Unlink(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-unlink", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m18.84 12.25 1.72-1.71h-.02a5.004 5.004 0 0 0-.12-7.07 5.006 5.006 0 0 0-6.95 0l-1.72 1.71"))),
		x.Child(x.Path(x.D("m5.17 11.75-1.71 1.71a5.004 5.004 0 0 0 .12 7.07 5.006 5.006 0 0 0 6.95 0l1.71-1.71"))),
		x.Child(x.Line(x.X1("8"), x.X2("8"), x.Y1("2"), x.Y2("5"))),
		x.Child(x.Line(x.X1("2"), x.X2("5"), x.Y1("8"), x.Y2("8"))),
		x.Child(x.Line(x.X1("16"), x.X2("16"), x.Y1("19"), x.Y2("22"))),
		x.Child(x.Line(x.X1("19"), x.X2("22"), x.Y1("16"), x.Y2("16"))),
	)
	return x.Svg(svgArgs...)
}
