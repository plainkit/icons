package lucide

import x "github.com/plainkit/blox"

// StarOff creates a Star Off Lucide icon.
func StarOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-star-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8.34 8.34 2 9.27l5 4.87L5.82 21 12 17.77 18.18 21l-.59-3.43"))),
		x.Child(x.Path(x.D("M18.42 12.76 22 9.27l-6.91-1L12 2l-1.44 2.91"))),
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("2"), x.Y2("22"))),
	)
	return x.Svg(svgArgs...)
}
