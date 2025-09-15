package lucide

import x "github.com/plainkit/html"

// AppWindowMac creates a App Window Mac Lucide icon.
func AppWindowMac(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-app-window-mac", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("16"), x.X("2"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M6 8h.01"))),
		x.Child(x.Path(x.D("M10 8h.01"))),
		x.Child(x.Path(x.D("M14 8h.01"))),
	)
	return x.Svg(svgArgs...)
}
