package lucide

import x "github.com/bloxui/blox"

// Blocks creates a Blocks Lucide icon.
func Blocks(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-blocks", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 22V7a1 1 0 0 0-1-1H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5a1 1 0 0 0-1-1H2"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("8"), x.X("14"), x.Y("2"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
