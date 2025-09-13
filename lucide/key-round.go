package lucide

import x "github.com/bloxui/blox"

// KeyRound creates a Key Round Lucide icon.
func KeyRound(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-key-round", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.586 17.414A2 2 0 0 0 2 18.828V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h.172a2 2 0 0 0 1.414-.586l.814-.814a6.5 6.5 0 1 0-4-4z"))),
		x.Child(x.Circle(x.Cx("16.5"), x.Cy("7.5"), x.R(".5"), x.Fill("currentColor"))),
	)
	return x.Svg(svgArgs...)
}
