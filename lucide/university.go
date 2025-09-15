package lucide

import x "github.com/plainkit/blox"

// University creates a University Lucide icon.
func University(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-university", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 21v-3a2 2 0 0 0-4 0v3"))),
		x.Child(x.Path(x.D("M18 12h.01"))),
		x.Child(x.Path(x.D("M18 16h.01"))),
		x.Child(x.Path(x.D("M22 7a1 1 0 0 0-1-1h-2a2 2 0 0 1-1.143-.359L13.143 2.36a2 2 0 0 0-2.286-.001L6.143 5.64A2 2 0 0 1 5 6H3a1 1 0 0 0-1 1v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2z"))),
		x.Child(x.Path(x.D("M6 12h.01"))),
		x.Child(x.Path(x.D("M6 16h.01"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("10"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
