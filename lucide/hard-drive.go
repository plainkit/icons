package lucide

import x "github.com/plainkit/html"

// HardDrive creates a Hard Drive Lucide icon.
func HardDrive(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-hard-drive", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("22"), x.X2("2"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Path(x.D("M5.45 5.11 2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11z"))),
		x.Child(x.Line(x.X1("6"), x.X2("6.01"), x.Y1("16"), x.Y2("16"))),
		x.Child(x.Line(x.X1("10"), x.X2("10.01"), x.Y1("16"), x.Y2("16"))),
	)
	return x.Svg(svgArgs...)
}
