package lucide

import x "github.com/plainkit/blox"

// MessageSquareDiff creates a Message Square Diff Lucide icon.
func MessageSquareDiff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-message-square-diff", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"))),
		x.Child(x.Path(x.D("M10 15h4"))),
		x.Child(x.Path(x.D("M10 9h4"))),
		x.Child(x.Path(x.D("M12 7v4"))),
	)
	return x.Svg(svgArgs...)
}
