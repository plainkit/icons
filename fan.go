package lucide

import x "github.com/bloxui/blox"

// Fan creates a Fan Lucide icon.
func Fan(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-fan", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.827 16.379a6.082 6.082 0 0 1-8.618-7.002l5.412 1.45a6.082 6.082 0 0 1 7.002-8.618l-1.45 5.412a6.082 6.082 0 0 1 8.618 7.002l-5.412-1.45a6.082 6.082 0 0 1-7.002 8.618l1.45-5.412Z"))),
		x.Child(x.Path(x.D("M12 12v.01"))),
	)
	return x.Svg(svgArgs...)
}