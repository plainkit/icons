package lucide

import x "github.com/bloxui/blox"

// ArrowUp01 creates a Arrow Up 0 1 Lucide icon.
func ArrowUp01(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-0-1", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 8 4-4 4 4"))),
		x.Child(x.Path(x.D("M7 4v16"))),
		x.Child(x.Rect(x.X("15"), x.Y("4"), x.RectWidth("4"), x.RectHeight("6"), x.Ry("2"))),
		x.Child(x.Path(x.D("M17 20v-6h-2"))),
		x.Child(x.Path(x.D("M15 20h4"))),
	)
	return x.Svg(svgArgs...)
}
