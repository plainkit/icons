package lucide

import x "github.com/bloxui/blox"

// ToggleRight creates a Toggle Right Lucide icon.
func ToggleRight(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-toggle-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("15"), x.Cy("12"), x.R("3"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("5"), x.Rx("7"))),
	)
	return x.Svg(svgArgs...)
}
