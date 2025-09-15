package lucide

import x "github.com/plainkit/html"

// ListRestart creates a List Restart Lucide icon.
func ListRestart(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-restart", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 5H3"))),
		x.Child(x.Path(x.D("M7 12H3"))),
		x.Child(x.Path(x.D("M7 19H3"))),
		x.Child(x.Path(x.D("M12 18a5 5 0 0 0 9-3 4.5 4.5 0 0 0-4.5-4.5c-1.33 0-2.54.54-3.41 1.41L11 14"))),
		x.Child(x.Path(x.D("M11 10v4h4"))),
	)
	return x.Svg(svgArgs...)
}
