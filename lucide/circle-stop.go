package lucide

import x "github.com/bloxui/blox"

// CircleStop creates a Circle Stop Lucide icon.
func CircleStop(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-stop", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("6"), x.X("9"), x.Y("9"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
