package lucide

import x "github.com/plainkit/blox"

// TriangleRight creates a Triangle Right Lucide icon.
func TriangleRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-triangle-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 18a2 2 0 0 1-2 2H3c-1.1 0-1.3-.6-.4-1.3L20.4 4.3c.9-.7 1.6-.4 1.6.7Z"))),
	)
	return x.Svg(svgArgs...)
}
