package lucide

import x "github.com/plainkit/blox"

// Plug creates a Plug Lucide icon.
func Plug(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-plug", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 22v-5"))),
		x.Child(x.Path(x.D("M9 8V2"))),
		x.Child(x.Path(x.D("M15 8V2"))),
		x.Child(x.Path(x.D("M18 8v5a4 4 0 0 1-4 4h-4a4 4 0 0 1-4-4V8Z"))),
	)
	return x.Svg(svgArgs...)
}
