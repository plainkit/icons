package lucide

import x "github.com/plainkit/html"

// GlobeLock creates a Globe Lock Lucide icon.
func GlobeLock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-globe-lock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15.686 15A14.5 14.5 0 0 1 12 22a14.5 14.5 0 0 1 0-20 10 10 0 1 0 9.542 13"))),
		x.Child(x.Path(x.D("M2 12h8.5"))),
		x.Child(x.Path(x.D("M20 6V4a2 2 0 1 0-4 0v2"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("5"), x.X("14"), x.Y("6"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
