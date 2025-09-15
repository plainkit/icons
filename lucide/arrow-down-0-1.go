package lucide

import x "github.com/plainkit/html"

// ArrowDown01 creates a Arrow Down 0 1 Lucide icon.
func ArrowDown01(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-0-1", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 16 4 4 4-4"))),
		x.Child(x.Path(x.D("M7 20V4"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("6"), x.X("15"), x.Y("4"), x.Ry("2"))),
		x.Child(x.Path(x.D("M17 20v-6h-2"))),
		x.Child(x.Path(x.D("M15 20h4"))),
	)
	return x.Svg(svgArgs...)
}
