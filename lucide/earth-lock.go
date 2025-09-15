package lucide

import x "github.com/plainkit/html"

// EarthLock creates a Earth Lock Lucide icon.
func EarthLock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-earth-lock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 3.34V5a3 3 0 0 0 3 3"))),
		x.Child(x.Path(x.D("M11 21.95V18a2 2 0 0 0-2-2 2 2 0 0 1-2-2v-1a2 2 0 0 0-2-2H2.05"))),
		x.Child(x.Path(x.D("M21.54 15H17a2 2 0 0 0-2 2v4.54"))),
		x.Child(x.Path(x.D("M12 2a10 10 0 1 0 9.54 13"))),
		x.Child(x.Path(x.D("M20 6V4a2 2 0 1 0-4 0v2"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("5"), x.X("14"), x.Y("6"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
