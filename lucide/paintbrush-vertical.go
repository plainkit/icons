package lucide

import x "github.com/plainkit/html"

// PaintbrushVertical creates a Paintbrush Vertical Lucide icon.
func PaintbrushVertical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-paintbrush-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 2v2"))),
		x.Child(x.Path(x.D("M14 2v4"))),
		x.Child(x.Path(x.D("M17 2a1 1 0 0 1 1 1v9H6V3a1 1 0 0 1 1-1z"))),
		x.Child(x.Path(x.D("M6 12a1 1 0 0 0-1 1v1a2 2 0 0 0 2 2h2a1 1 0 0 1 1 1v2.9a2 2 0 1 0 4 0V17a1 1 0 0 1 1-1h2a2 2 0 0 0 2-2v-1a1 1 0 0 0-1-1"))),
	)
	return x.Svg(svgArgs...)
}
