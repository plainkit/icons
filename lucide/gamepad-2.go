package lucide

import x "github.com/plainkit/html"

// Gamepad2 creates a Gamepad 2 Lucide icon.
func Gamepad2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-gamepad-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("6"), x.X2("10"), x.Y1("11"), x.Y2("11"))),
		x.Child(x.Line(x.X1("8"), x.X2("8"), x.Y1("9"), x.Y2("13"))),
		x.Child(x.Line(x.X1("15"), x.X2("15.01"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("18"), x.X2("18.01"), x.Y1("10"), x.Y2("10"))),
		x.Child(x.Path(x.D("M17.32 5H6.68a4 4 0 0 0-3.978 3.59c-.006.052-.01.101-.017.152C2.604 9.416 2 14.456 2 16a3 3 0 0 0 3 3c1 0 1.5-.5 2-1l1.414-1.414A2 2 0 0 1 9.828 16h4.344a2 2 0 0 1 1.414.586L17 18c.5.5 1 1 2 1a3 3 0 0 0 3-3c0-1.545-.604-6.584-.685-7.258-.007-.05-.011-.1-.017-.151A4 4 0 0 0 17.32 5z"))),
	)
	return x.Svg(svgArgs...)
}
