package lucide

import x "github.com/bloxui/blox"

// PictureInPicture2 creates a Picture In Picture 2 Lucide icon.
func PictureInPicture2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-picture-in-picture-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 9V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10c0 1.1.9 2 2 2h4"))),
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("7"), x.X("12"), x.Y("13"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
