package lucide

import x "github.com/bloxui/blox"

// MapPinned creates a Map Pinned Lucide icon.
func MapPinned(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-map-pinned", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 8c0 3.613-3.869 7.429-5.393 8.795a1 1 0 0 1-1.214 0C9.87 15.429 6 11.613 6 8a6 6 0 0 1 12 0"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("8"), x.R("2"))),
		x.Child(x.Path(x.D("M8.714 14h-3.71a1 1 0 0 0-.948.683l-2.004 6A1 1 0 0 0 3 22h18a1 1 0 0 0 .948-1.316l-2-6a1 1 0 0 0-.949-.684h-3.712"))),
	)
	return x.Svg(svgArgs...)
}
