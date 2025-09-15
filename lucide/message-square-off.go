package lucide

import x "github.com/plainkit/blox"

// MessageSquareOff creates a Message Square Off Lucide icon.
func MessageSquareOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-message-square-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19 19H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.7.7 0 0 1 2 21.286V5a2 2 0 0 1 1.184-1.826"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M8.656 3H20a2 2 0 0 1 2 2v11.344"))),
	)
	return x.Svg(svgArgs...)
}
