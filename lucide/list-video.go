package lucide

import x "github.com/bloxui/blox"

// ListVideo creates a List Video Lucide icon.
func ListVideo(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-list-video", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 5H3"))),
		x.Child(x.Path(x.D("M10 12H3"))),
		x.Child(x.Path(x.D("M10 19H3"))),
		x.Child(x.Path(x.D("M15 12.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997a1 1 0 0 1-1.517-.86z"))),
	)
	return x.Svg(svgArgs...)
}
