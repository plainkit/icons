package lucide

import x "github.com/bloxui/blox"

// Container creates a Container Lucide icon.
func Container(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-container", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 7.7c0-.6-.4-1.2-.8-1.5l-6.3-3.9a1.72 1.72 0 0 0-1.7 0l-10.3 6c-.5.2-.9.8-.9 1.4v6.6c0 .5.4 1.2.8 1.5l6.3 3.9a1.72 1.72 0 0 0 1.7 0l10.3-6c.5-.3.9-1 .9-1.5Z"))),
		x.Child(x.Path(x.D("M10 21.9V14L2.1 9.1"))),
		x.Child(x.Path(x.D("m10 14 11.9-6.9"))),
		x.Child(x.Path(x.D("M14 19.8v-8.1"))),
		x.Child(x.Path(x.D("M18 17.5V9.4"))),
	)
	return x.Svg(svgArgs...)
}