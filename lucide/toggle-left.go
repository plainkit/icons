package lucide

import x "github.com/plainkit/blox"

// ToggleLeft creates a Toggle Left Lucide icon.
func ToggleLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-toggle-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("9"), x.Cy("12"), x.R("3"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("5"), x.Rx("7"))),
	)
	return x.Svg(svgArgs...)
}
