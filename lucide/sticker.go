package lucide

import x "github.com/bloxui/blox"

// Sticker creates a Sticker Lucide icon.
func Sticker(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-sticker", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15.5 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h14a2 2 0 0 0 2-2V8.5L15.5 3Z"))),
		x.Child(x.Path(x.D("M14 3v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M8 13h.01"))),
		x.Child(x.Path(x.D("M16 13h.01"))),
		x.Child(x.Path(x.D("M10 16s.8 1 2 1c1.3 0 2-1 2-1"))),
	)
	return x.Svg(svgArgs...)
}
