package generator

import (
	"fmt"
	"math"
)

// Osc is an oscillator
type Osc struct {
	Shape     WaveType
	Amplitude float64
	DcOffset  float64
	Freq      float64
	// SampleRate
	Fs                int
	PhaseOffset       float64
	CurrentPhaseAngle float64
	phaseAngleIncr    float64
}

// NewOsc returns a new oscillator, note that if you change the phase offset of the returned osc,
// you also need to set the CurrentPhaseAngle
func NewOsc(shape WaveType, hz float64, fs int) *Osc {
	return &Osc{Shape: shape, Amplitude: 1, Freq: hz, Fs: fs, phaseAngleIncr: ((hz * TwoPi) / float64(fs))}
}

// Signal uses the osc to generate a discreet signal
func (o *Osc) Signal(length int) []float64 {
	output := make([]float64, length)
	for i := 0; i < length; i++ {
		output[i] = o.Sample()
	}
	return output
}

// Sample returns the next sample generated by the oscillator
func (o *Osc) Sample() (output float64) {
	if o == nil {
		return
	}

	if o.CurrentPhaseAngle < -math.Pi {
		o.CurrentPhaseAngle += TwoPi
	} else if o.CurrentPhaseAngle > math.Pi {
		o.CurrentPhaseAngle -= TwoPi
	}

	switch o.Shape {
	case WaveSine:
		output = o.Amplitude*Sine(o.CurrentPhaseAngle) + o.DcOffset
	case WaveTriangle:
		output = o.Amplitude*Triangle(o.CurrentPhaseAngle) + o.DcOffset
	case WaveSaw:
		output = o.Amplitude*Sawtooth(o.CurrentPhaseAngle) + o.DcOffset
	case WaveSqr:
		fmt.Println(o.CurrentPhaseAngle)
		output = o.Amplitude*Square(o.CurrentPhaseAngle) + o.DcOffset
	}

	o.CurrentPhaseAngle += o.phaseAngleIncr
	return output
}
