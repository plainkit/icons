package lucide

import x "github.com/bloxui/blox"

// BottleWine creates a Bottle Wine Lucide icon.
func BottleWine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bottle-wine", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a6 6 0 0 0 1.2 3.6l.6.8A6 6 0 0 1 17 13v8a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-8a6 6 0 0 1 1.2-3.6l.6-.8A6 6 0 0 0 10 5z"))),
		x.Child(x.Path(x.D("M17 13h-4a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h4"))),
	)
	return x.Svg(svgArgs...)
}
