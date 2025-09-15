package lucide

import x "github.com/plainkit/html"

// DiscAlbum creates a Disc Album Lucide icon.
func DiscAlbum(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-disc-album", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("5"))),
		x.Child(x.Path(x.D("M12 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
