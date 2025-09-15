package lucide

import x "github.com/plainkit/html"

// ArrowDown10 creates a Arrow Down 1 0 Lucide icon.
func ArrowDown10(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-1-0", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 16 4 4 4-4"))),
		x.Child(x.Path(x.D("M7 20V4"))),
		x.Child(x.Path(x.D("M17 10V4h-2"))),
		x.Child(x.Path(x.D("M15 10h4"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("6"), x.X("15"), x.Y("14"), x.Ry("2"))),
	)
	return x.Svg(svgArgs...)
}
