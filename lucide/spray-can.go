package lucide

import x "github.com/plainkit/blox"

// SprayCan creates a Spray Can Lucide icon.
func SprayCan(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-spray-can", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3h.01"))),
		x.Child(x.Path(x.D("M7 5h.01"))),
		x.Child(x.Path(x.D("M11 7h.01"))),
		x.Child(x.Path(x.D("M3 7h.01"))),
		x.Child(x.Path(x.D("M7 9h.01"))),
		x.Child(x.Path(x.D("M3 11h.01"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("4"), x.X("15"), x.Y("5"))),
		x.Child(x.Path(x.D("m19 9 2 2v10c0 .6-.4 1-1 1h-6c-.6 0-1-.4-1-1V11l2-2"))),
		x.Child(x.Path(x.D("m13 14 8-2"))),
		x.Child(x.Path(x.D("m13 19 8-2"))),
	)
	return x.Svg(svgArgs...)
}
