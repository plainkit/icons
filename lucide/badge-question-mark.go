package lucide

import x "github.com/plainkit/html"

// BadgeQuestionMark creates a Badge Question Mark Lucide icon.
func BadgeQuestionMark(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-badge-question-mark", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		x.Child(x.Path(x.D("M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"))),
		x.Child(x.Line(x.X1("12"), x.X2("12.01"), x.Y1("17"), x.Y2("17"))),
	)
	return x.Svg(svgArgs...)
}
