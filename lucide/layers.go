package lucide

import x "github.com/bloxui/blox"

// Layers creates a Layers Lucide icon.
func Layers(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-layers", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.83 2.18a2 2 0 0 0-1.66 0L2.6 6.08a1 1 0 0 0 0 1.83l8.58 3.91a2 2 0 0 0 1.66 0l8.58-3.9a1 1 0 0 0 0-1.83z"))),
		x.Child(x.Path(x.D("M2 12a1 1 0 0 0 .58.91l8.6 3.91a2 2 0 0 0 1.65 0l8.58-3.9A1 1 0 0 0 22 12"))),
		x.Child(x.Path(x.D("M2 17a1 1 0 0 0 .58.91l8.6 3.91a2 2 0 0 0 1.65 0l8.58-3.9A1 1 0 0 0 22 17"))),
	)
	return x.Svg(svgArgs...)
}
