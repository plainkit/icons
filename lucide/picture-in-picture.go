package lucide

import x "github.com/bloxui/blox"

// PictureInPicture creates a Picture In Picture Lucide icon.
func PictureInPicture(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-picture-in-picture", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 10h6V4"))),
		x.Child(x.Path(x.D("m2 4 6 6"))),
		x.Child(x.Path(x.D("M21 10V7a2 2 0 0 0-2-2h-7"))),
		x.Child(x.Path(x.D("M3 14v2a2 2 0 0 0 2 2h3"))),
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("7"), x.X("12"), x.Y("14"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
