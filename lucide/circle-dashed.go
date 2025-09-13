package lucide

import x "github.com/bloxui/blox"

// CircleDashed creates a Circle Dashed Lucide icon.
func CircleDashed(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-circle-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.1 2.182a10 10 0 0 1 3.8 0"))),
		x.Child(x.Path(x.D("M13.9 21.818a10 10 0 0 1-3.8 0"))),
		x.Child(x.Path(x.D("M17.609 3.721a10 10 0 0 1 2.69 2.7"))),
		x.Child(x.Path(x.D("M2.182 13.9a10 10 0 0 1 0-3.8"))),
		x.Child(x.Path(x.D("M20.279 17.609a10 10 0 0 1-2.7 2.69"))),
		x.Child(x.Path(x.D("M21.818 10.1a10 10 0 0 1 0 3.8"))),
		x.Child(x.Path(x.D("M3.721 6.391a10 10 0 0 1 2.7-2.69"))),
		x.Child(x.Path(x.D("M6.391 20.279a10 10 0 0 1-2.69-2.7"))),
	)
	return x.Svg(svgArgs...)
}
