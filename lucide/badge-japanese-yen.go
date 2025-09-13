package lucide

import x "github.com/bloxui/blox"

// BadgeJapaneseYen creates a Badge Japanese Yen Lucide icon.
func BadgeJapaneseYen(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-badge-japanese-yen", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		x.Child(x.Path(x.D("m9 8 3 3v7"))),
		x.Child(x.Path(x.D("m12 11 3-3"))),
		x.Child(x.Path(x.D("M9 12h6"))),
		x.Child(x.Path(x.D("M9 16h6"))),
	)
	return x.Svg(svgArgs...)
}
