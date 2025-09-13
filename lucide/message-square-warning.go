package lucide

import x "github.com/bloxui/blox"

// MessageSquareWarning creates a Message Square Warning Lucide icon.
func MessageSquareWarning(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-message-square-warning", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"))),
		x.Child(x.Path(x.D("M12 15h.01"))),
		x.Child(x.Path(x.D("M12 7v4"))),
	)
	return x.Svg(svgArgs...)
}
