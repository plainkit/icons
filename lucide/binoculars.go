package lucide

import x "github.com/bloxui/blox"

// Binoculars creates a Binoculars Lucide icon.
func Binoculars(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-binoculars", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 10h4"))),
		x.Child(x.Path(x.D("M19 7V4a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v3"))),
		x.Child(x.Path(x.D("M20 21a2 2 0 0 0 2-2v-3.851c0-1.39-2-2.962-2-4.829V8a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v11a2 2 0 0 0 2 2z"))),
		x.Child(x.Path(x.D("M 22 16 L 2 16"))),
		x.Child(x.Path(x.D("M4 21a2 2 0 0 1-2-2v-3.851c0-1.39 2-2.962 2-4.829V8a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v11a2 2 0 0 1-2 2z"))),
		x.Child(x.Path(x.D("M9 7V4a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v3"))),
	)
	return x.Svg(svgArgs...)
}
