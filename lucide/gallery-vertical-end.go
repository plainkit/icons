package lucide

import x "github.com/plainkit/html"

// GalleryVerticalEnd creates a Gallery Vertical End Lucide icon.
func GalleryVerticalEnd(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-gallery-vertical-end", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 2h10"))),
		x.Child(x.Path(x.D("M5 6h14"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("12"), x.X("3"), x.Y("10"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
