package lucide

import x "github.com/plainkit/blox"

// Ship creates a Ship Lucide icon.
func Ship(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ship", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 10.189V14"))),
		x.Child(x.Path(x.D("M12 2v3"))),
		x.Child(x.Path(x.D("M19 13V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v6"))),
		x.Child(x.Path(x.D("M19.38 20A11.6 11.6 0 0 0 21 14l-8.188-3.639a2 2 0 0 0-1.624 0L3 14a11.6 11.6 0 0 0 2.81 7.76"))),
		x.Child(x.Path(x.D("M2 21c.6.5 1.2 1 2.5 1 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1s1.2 1 2.5 1c2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"))),
	)
	return x.Svg(svgArgs...)
}
