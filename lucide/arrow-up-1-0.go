package lucide

import x "github.com/bloxui/blox"

// ArrowUp10 creates a Arrow Up 1 0 Lucide icon.
func ArrowUp10(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-1-0", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 8 4-4 4 4"))),
		x.Child(x.Path(x.D("M7 4v16"))),
		x.Child(x.Path(x.D("M17 10V4h-2"))),
		x.Child(x.Path(x.D("M15 10h4"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("6"), x.X("15"), x.Y("14"), x.Ry("2"))),
	)
	return x.Svg(svgArgs...)
}
