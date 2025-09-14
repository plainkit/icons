package lucide

import x "github.com/bloxui/blox"

// RectangleGoggles creates a Rectangle Goggles Lucide icon.
func RectangleGoggles(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rectangle-goggles", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 6a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4a2 2 0 0 1-1.6-.8l-1.6-2.13a1 1 0 0 0-1.6 0L9.6 17.2A2 2 0 0 1 8 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2z"))),
	)
	return x.Svg(svgArgs...)
}
