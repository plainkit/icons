package lucide

import x "github.com/plainkit/html"

// ArchiveX creates a Archive X Lucide icon.
func ArchiveX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-archive-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("5"), x.X("2"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Path(x.D("M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8"))),
		x.Child(x.Path(x.D("m9.5 17 5-5"))),
		x.Child(x.Path(x.D("m9.5 12 5 5"))),
	)
	return x.Svg(svgArgs...)
}
