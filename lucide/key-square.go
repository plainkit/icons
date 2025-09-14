package lucide

import x "github.com/bloxui/blox"

// KeySquare creates a Key Square Lucide icon.
func KeySquare(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-key-square", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.4 2.7a2.5 2.5 0 0 1 3.4 0l5.5 5.5a2.5 2.5 0 0 1 0 3.4l-3.7 3.7a2.5 2.5 0 0 1-3.4 0L8.7 9.8a2.5 2.5 0 0 1 0-3.4z"))),
		x.Child(x.Path(x.D("m14 7 3 3"))),
		x.Child(x.Path(x.D("m9.4 10.6-6.814 6.814A2 2 0 0 0 2 18.828V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h.172a2 2 0 0 0 1.414-.586l.814-.814"))),
	)
	return x.Svg(svgArgs...)
}
