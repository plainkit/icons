package lucide

import x "github.com/plainkit/blox"

// GalleryThumbnails creates a Gallery Thumbnails Lucide icon.
func GalleryThumbnails(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-gallery-thumbnails", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("14"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M4 21h1"))),
		x.Child(x.Path(x.D("M9 21h1"))),
		x.Child(x.Path(x.D("M14 21h1"))),
		x.Child(x.Path(x.D("M19 21h1"))),
	)
	return x.Svg(svgArgs...)
}
