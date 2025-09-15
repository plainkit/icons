package lucide

import x "github.com/plainkit/html"

// VibrateOff creates a Vibrate Off Lucide icon.
func VibrateOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-vibrate-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m2 8 2 2-2 2 2 2-2 2"))),
		x.Child(x.Path(x.D("m22 8-2 2 2 2-2 2 2 2"))),
		x.Child(x.Path(x.D("M8 8v10c0 .55.45 1 1 1h6c.55 0 1-.45 1-1v-2"))),
		x.Child(x.Path(x.D("M16 10.34V6c0-.55-.45-1-1-1h-4.34"))),
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("2"), x.Y2("22"))),
	)
	return x.Svg(svgArgs...)
}
