package lucide

import x "github.com/plainkit/blox"

// TvMinimal creates a Tv Minimal Lucide icon.
func TvMinimal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tv-minimal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 21h10"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
