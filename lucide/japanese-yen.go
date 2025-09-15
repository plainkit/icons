package lucide

import x "github.com/plainkit/blox"

// JapaneseYen creates a Japanese Yen Lucide icon.
func JapaneseYen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-japanese-yen", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 9.5V21m0-11.5L6 3m6 6.5L18 3"))),
		x.Child(x.Path(x.D("M6 15h12"))),
		x.Child(x.Path(x.D("M6 11h12"))),
	)
	return x.Svg(svgArgs...)
}
