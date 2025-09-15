package lucide

import x "github.com/plainkit/blox"

// GalleryVertical creates a Gallery Vertical Lucide icon.
func GalleryVertical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-gallery-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 2h18"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("12"), x.X("3"), x.Y("6"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 22h18"))),
	)
	return x.Svg(svgArgs...)
}
