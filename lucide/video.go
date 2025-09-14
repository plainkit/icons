package lucide

import x "github.com/bloxui/blox"

// Video creates a Video Lucide icon.
func Video(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-video", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 13 5.223 3.482a.5.5 0 0 0 .777-.416V7.87a.5.5 0 0 0-.752-.432L16 10.5"))),
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("12"), x.X("2"), x.Y("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
