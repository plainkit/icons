package lucide

import x "github.com/bloxui/blox"

// Twitch creates a Twitch Lucide icon.
func Twitch(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-twitch", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 2H3v16h5v4l4-4h5l4-4V2zm-10 9V7m5 4V7"))),
	)
	return x.Svg(svgArgs...)
}
