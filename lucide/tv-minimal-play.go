package lucide

import x "github.com/bloxui/blox"

// TvMinimalPlay creates a Tv Minimal Play Lucide icon.
func TvMinimalPlay(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tv-minimal-play", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15.033 9.44a.647.647 0 0 1 0 1.12l-4.065 2.352a.645.645 0 0 1-.968-.56V7.648a.645.645 0 0 1 .967-.56z"))),
		x.Child(x.Path(x.D("M7 21h10"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
