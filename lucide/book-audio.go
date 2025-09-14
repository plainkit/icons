package lucide

import x "github.com/bloxui/blox"

// BookAudio creates a Book Audio Lucide icon.
func BookAudio(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-book-audio", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v7"))),
		x.Child(x.Path(x.D("M16 8v3"))),
		x.Child(x.Path(x.D("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		x.Child(x.Path(x.D("M8 8v3"))),
	)
	return x.Svg(svgArgs...)
}
