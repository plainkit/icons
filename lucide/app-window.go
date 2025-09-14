package lucide

import x "github.com/bloxui/blox"

// AppWindow creates a App Window Lucide icon.
func AppWindow(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-app-window", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("16"), x.X("2"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M10 4v4"))),
		x.Child(x.Path(x.D("M2 8h20"))),
		x.Child(x.Path(x.D("M6 4v4"))),
	)
	return x.Svg(svgArgs...)
}
