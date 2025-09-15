package lucide

import x "github.com/plainkit/html"

// Mailbox creates a Mailbox Lucide icon.
func Mailbox(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mailbox", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9.5C2 7 4 5 6.5 5H18c2.2 0 4 1.8 4 4v8Z"))),
		x.Child(x.Polyline(x.Points("15,9 18,9 18,11"))),
		x.Child(x.Path(x.D("M6.5 5C9 5 11 7 11 9.5V17a2 2 0 0 1-2 2"))),
		x.Child(x.Line(x.X1("6"), x.X2("7"), x.Y1("10"), x.Y2("10"))),
	)
	return x.Svg(svgArgs...)
}
