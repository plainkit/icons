package lucide

import x "github.com/bloxui/blox"

// Plug2 creates a Plug 2 Lucide icon.
func Plug2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-plug-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 2v6"))),
		x.Child(x.Path(x.D("M15 2v6"))),
		x.Child(x.Path(x.D("M12 17v5"))),
		x.Child(x.Path(x.D("M5 8h14"))),
		x.Child(x.Path(x.D("M6 11V8h12v3a6 6 0 1 1-12 0Z"))),
	)
	return x.Svg(svgArgs...)
}
