package lucide

import x "github.com/plainkit/blox"

// Lock creates a Lock Lucide icon.
func Lock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-lock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("11"), x.X("3"), x.Y("11"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("M7 11V7a5 5 0 0 1 10 0v4"))),
	)
	return x.Svg(svgArgs...)
}
