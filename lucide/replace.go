package lucide

import x "github.com/plainkit/html"

// Replace creates a Replace Lucide icon.
func Replace(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-replace", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 4a2 2 0 0 1 2-2"))),
		x.Child(x.Path(x.D("M16 10a2 2 0 0 1-2-2"))),
		x.Child(x.Path(x.D("M20 2a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M22 8a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("m3 7 3 3 3-3"))),
		x.Child(x.Path(x.D("M6 10V5a3 3 0 0 1 3-3h1"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("8"), x.X("2"), x.Y("14"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
