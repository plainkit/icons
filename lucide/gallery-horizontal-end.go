package lucide

import x "github.com/bloxui/blox"

// GalleryHorizontalEnd creates a Gallery Horizontal End Lucide icon.
func GalleryHorizontalEnd(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-gallery-horizontal-end", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 7v10"))),
		x.Child(x.Path(x.D("M6 5v14"))),
		x.Child(x.Rect(x.RectWidth("12"), x.RectHeight("18"), x.X("10"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
