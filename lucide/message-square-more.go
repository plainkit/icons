package lucide

import x "github.com/plainkit/blox"

// MessageSquareMore creates a Message Square More Lucide icon.
func MessageSquareMore(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-message-square-more", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"))),
		x.Child(x.Path(x.D("M12 11h.01"))),
		x.Child(x.Path(x.D("M16 11h.01"))),
		x.Child(x.Path(x.D("M8 11h.01"))),
	)
	return x.Svg(svgArgs...)
}
