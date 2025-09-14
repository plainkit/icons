package lucide

import x "github.com/bloxui/blox"

// BrushCleaning creates a Brush Cleaning Lucide icon.
func BrushCleaning(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-brush-cleaning", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 22-1-4"))),
		x.Child(x.Path(x.D("M19 13.99a1 1 0 0 0 1-1V12a2 2 0 0 0-2-2h-3a1 1 0 0 1-1-1V4a2 2 0 0 0-4 0v5a1 1 0 0 1-1 1H6a2 2 0 0 0-2 2v.99a1 1 0 0 0 1 1"))),
		x.Child(x.Path(x.D("M5 14h14l1.973 6.767A1 1 0 0 1 20 22H4a1 1 0 0 1-.973-1.233z"))),
		x.Child(x.Path(x.D("m8 22 1-4"))),
	)
	return x.Svg(svgArgs...)
}
