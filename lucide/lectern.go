package lucide

import x "github.com/plainkit/html"

// Lectern creates a Lectern Lucide icon.
func Lectern(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-lectern", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 12h3a2 2 0 0 0 1.902-1.38l1.056-3.333A1 1 0 0 0 21 6H3a1 1 0 0 0-.958 1.287l1.056 3.334A2 2 0 0 0 5 12h3"))),
		x.Child(x.Path(x.D("M18 6V3a1 1 0 0 0-1-1h-3"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("12"), x.X("8"), x.Y("10"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
