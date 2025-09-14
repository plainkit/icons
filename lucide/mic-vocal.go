package lucide

import x "github.com/bloxui/blox"

// MicVocal creates a Mic Vocal Lucide icon.
func MicVocal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mic-vocal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m11 7.601-5.994 8.19a1 1 0 0 0 .1 1.298l.817.818a1 1 0 0 0 1.314.087L15.09 12"))),
		x.Child(x.Path(x.D("M16.5 21.174C15.5 20.5 14.372 20 13 20c-2.058 0-3.928 2.356-6 2-2.072-.356-2.775-3.369-1.5-4.5"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("7"), x.R("5"))),
	)
	return x.Svg(svgArgs...)
}
