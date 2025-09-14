package lucide

import x "github.com/bloxui/blox"

// WavesLadder creates a Waves Ladder Lucide icon.
func WavesLadder(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-waves-ladder", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19 5a2 2 0 0 0-2 2v11"))),
		x.Child(x.Path(x.D("M2 18c.6.5 1.2 1 2.5 1 2.5 0 2.5-2 5-2 2.6 0 2.4 2 5 2 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"))),
		x.Child(x.Path(x.D("M7 13h10"))),
		x.Child(x.Path(x.D("M7 9h10"))),
		x.Child(x.Path(x.D("M9 5a2 2 0 0 0-2 2v11"))),
	)
	return x.Svg(svgArgs...)
}
