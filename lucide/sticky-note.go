package lucide

import x "github.com/plainkit/html"

// StickyNote creates a Sticky Note Lucide icon.
func StickyNote(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sticky-note", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8Z"))),
		x.Child(x.Path(x.D("M15 3v4a2 2 0 0 0 2 2h4"))),
	)
	return x.Svg(svgArgs...)
}
