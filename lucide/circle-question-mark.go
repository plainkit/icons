package lucide

import x "github.com/plainkit/blox"

// CircleQuestionMark creates a Circle Question Mark Lucide icon.
func CircleQuestionMark(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-question-mark", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"))),
		x.Child(x.Path(x.D("M12 17h.01"))),
	)
	return x.Svg(svgArgs...)
}
