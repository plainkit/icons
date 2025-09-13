package lucide

import x "github.com/bloxui/blox"

// AudioWaveform creates a Audio Waveform Lucide icon.
func AudioWaveform(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-audio-waveform", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 13a2 2 0 0 0 2-2V7a2 2 0 0 1 4 0v13a2 2 0 0 0 4 0V4a2 2 0 0 1 4 0v13a2 2 0 0 0 4 0v-4a2 2 0 0 1 2-2"))),
	)
	return x.Svg(svgArgs...)
}
