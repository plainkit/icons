package lucide

import x "github.com/bloxui/blox"

// GalleryHorizontal creates a Gallery Horizontal Lucide icon.
func GalleryHorizontal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-gallery-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 3v18"))),
		x.Child(x.Rect(x.RectWidth("12"), x.RectHeight("18"), x.X("6"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M22 3v18"))),
	)
	return x.Svg(svgArgs...)
}
