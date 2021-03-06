package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// Haptic effects.
// (https://wiki.libsdl.org/SDL_HapticEffect)
const (
	HAPTIC_CONSTANT     = C.SDL_HAPTIC_CONSTANT     // constant haptic effect
	HAPTIC_SINE         = C.SDL_HAPTIC_SINE         // periodic haptic effect that simulates sine waves
	HAPTIC_LEFTRIGHT    = C.SDL_HAPTIC_LEFTRIGHT    // haptic effect for direct control over high/low frequency motors
	HAPTIC_TRIANGLE     = C.SDL_HAPTIC_TRIANGLE     // periodic haptic effect that simulates triangular waves
	HAPTIC_SAWTOOTHUP   = C.SDL_HAPTIC_SAWTOOTHUP   // periodic haptic effect that simulates saw tooth up waves
	HAPTIC_SAWTOOTHDOWN = C.SDL_HAPTIC_SAWTOOTHDOWN // periodic haptic effect that simulates saw tooth down waves
	HAPTIC_RAMP         = C.SDL_HAPTIC_RAMP         // ramp haptic effect
	HAPTIC_SPRING       = C.SDL_HAPTIC_SPRING       // condition haptic effect that simulates a spring.  Effect is based on the axes position
	HAPTIC_DAMPER       = C.SDL_HAPTIC_DAMPER       // condition haptic effect that simulates dampening.  Effect is based on the axes velocity
	HAPTIC_INERTIA      = C.SDL_HAPTIC_INERTIA      // condition haptic effect that simulates inertia.  Effect is based on the axes acceleration
	HAPTIC_FRICTION     = C.SDL_HAPTIC_FRICTION     // condition haptic effect that simulates friction.  Effect is based on the axes movement
	HAPTIC_CUSTOM       = C.SDL_HAPTIC_CUSTOM       // user defined custom haptic effect
	HAPTIC_GAIN         = C.SDL_HAPTIC_GAIN         // device supports setting the global gain
	HAPTIC_AUTOCENTER   = C.SDL_HAPTIC_AUTOCENTER   // device supports setting autocenter
	HAPTIC_STATUS       = C.SDL_HAPTIC_STATUS       // device can be queried for effect status
	HAPTIC_PAUSE        = C.SDL_HAPTIC_PAUSE        // device can be paused
	//HAPTIC_SQUARE = C.SDL_HAPTIC_SQUARE (back in SDL 2.1)
)

// Direction encodings.
// (https://wiki.libsdl.org/SDL_HapticDirection)
const (
	HAPTIC_POLAR     = C.SDL_HAPTIC_POLAR     // uses polar coordinates for the direction
	HAPTIC_CARTESIAN = C.SDL_HAPTIC_CARTESIAN // uses cartesian coordinates for the direction
	HAPTIC_SPHERICAL = C.SDL_HAPTIC_SPHERICAL // uses spherical coordinates for the direction
	HAPTIC_INFINITY  = C.SDL_HAPTIC_INFINITY  // used to play a device an infinite number of times
)

// Haptic identifies an SDL haptic.
// (https://wiki.libsdl.org/CategoryForceFeedback)
type Haptic C.SDL_Haptic

// HapticDirection contains a haptic direction.
// (https://wiki.libsdl.org/SDL_HapticDirection)
type HapticDirection struct {
	Type byte     // the type of encoding
	dir  [3]int32 // the encoded direction
}

// HapticConstant contains a template for a constant effect.
// (https://wiki.libsdl.org/SDL_HapticConstant)
type HapticConstant struct {
	Type         uint16          // HAPTIC_CONSTANT
	Direction    HapticDirection // direction of the effect
	Length       uint32          // duration of the effect
	Delay        uint16          // delay before starting the effect
	Button       uint16          // button that triggers the effect
	Interval     uint16          // how soon it can be triggered again after button
	Level        int16           // strength of the constant effect
	AttackLength uint16          // duration of the attack
	AttackLevel  uint16          // level at the start of the attack
	FadeLength   uint16          // duration of the fade
	FadeLevel    uint16          // level at the end of the fade
}

// HapticPeriodic contains a template for a periodic effect.
// (https://wiki.libsdl.org/SDL_HapticPeriodic)
type HapticPeriodic struct {
	Type         uint16          // HAPTIC_SINE, HAPTIC_LEFTRIGHT, HAPTIC_TRIANGLE, HAPTIC_SAWTOOTHUP, HAPTIC_SAWTOOTHDOWN
	Direction    HapticDirection // direction of the effect
	Length       uint32          // duration of the effect
	Delay        uint16          // delay before starting the effect
	Button       uint16          // button that triggers the effect
	Interval     uint16          // how soon it can be triggered again after button
	Period       uint16          // period of the wave
	Magnitude    int16           // peak value; if negative, equivalent to 180 degrees extra phase shift
	Offset       int16           // mean value of the wave
	Phase        uint16          // positive phase shift given by hundredth of a degree
	AttackLength uint16          // duration of the attack
	AttackLevel  uint16          // level at the start of the attack
	FadeLength   uint16          // duration of the fade
	FadeLevel    uint16          // level at the end of the fade
}

// HapticCondition contains a template for a condition effect.
// (https://wiki.libsdl.org/SDL_HapticCondition)
type HapticCondition struct {
	Type       uint16          // HAPTIC_SPRING, HAPTIC_DAMPER, HAPTIC_INERTIA, HAPTIC_FRICTION
	Direction  HapticDirection // direction of the effect - not used at the moment
	Length     uint32          // duration of the effect
	Delay      uint16          // delay before starting the effect
	Button     uint16          // button that triggers the effect
	Interval   uint16          // how soon it can be triggered again after button
	RightSat   [3]uint16       // level when joystick is to the positive side; max 0xFFFF
	LeftSat    [3]uint16       // level when joystick is to the negative side; max 0xFFFF
	RightCoeff [3]int16        // how fast to increase the force towards the positive side
	LeftCoeff  [3]int16        // how fast to increase the force towards the negative side
	Deadband   [3]uint16       // size of the dead zone; max 0xFFFF: whole axis-range when 0-centered
	Center     [3]int16        // position of the dead zone
}

// HapticRamp contains a template for a ramp effect.
// (https://wiki.libsdl.org/SDL_HapticRamp)
type HapticRamp struct {
	Type         uint16          // HAPTIC_RAMP
	Direction    HapticDirection // direction of the effect
	Length       uint32          // duration of the effect
	Delay        uint16          // delay before starting the effect
	Button       uint16          // button that triggers the effect
	Interval     uint16          // how soon it can be triggered again after button
	Start        int16           // beginning strength level
	End          int16           // ending strength level
	AttackLength uint16          // duration of the attack
	AttackLevel  uint16          // level at the start of the attack
	FadeLength   uint16          // duration of the fade
	FadeLevel    uint16          // level at the end of the fade
}

// HapticLeftRight contains a template for a left/right effect.
// (https://wiki.libsdl.org/SDL_HapticLeftRight)
type HapticLeftRight struct {
	Type           uint16 // HAPTIC_LEFTRIGHT
	Length         uint32 // duration of the effect
	LargeMagnitude uint16 // control of the large controller motor
	SmallMagnitude uint16 // control of the small controller motor
}

// HapticCustom contains a template for a custom effect.
// (https://wiki.libsdl.org/SDL_HapticCustom)
type HapticCustom struct {
	Type         uint16          // SDL_HAPTIC_CUSTOM
	Direction    HapticDirection // direction of the effect
	Length       uint32          // duration of the effect
	Delay        uint16          // delay before starting the effect
	Button       uint16          // button that triggers the effect
	Interval     uint16          // how soon it can be triggered again after button
	Channels     uint8           // axes to use, minimum of 1
	Period       uint16          // sample periods
	Samples      uint16          // amount of samples
	Data         *uint16         // should contain channels*samples items
	AttackLength uint16          // duration of the attack
	AttackLevel  uint16          // level at the start of the attack
	FadeLength   uint16          // duration of the fade
	FadeLevel    uint16          // level at the end of the fade
}

// HapticEffect union that contains the generic template for any haptic effect.
// (https://wiki.libsdl.org/SDL_HapticEffect)
type HapticEffect C.SDL_HapticEffect

// Type returns the effect type.
// (https://wiki.libsdl.org/SDL_HapticEffect)
func (he HapticEffect) Type() uint16 {
	return *((*uint16)(unsafe.Pointer(&he[0])))
}

// Constant returns the constant effect.
// (https://wiki.libsdl.org/SDL_HapticConstant)
func (he HapticEffect) Constant() *HapticConstant {
	return (*HapticConstant)(unsafe.Pointer(&he[0]))
}

// Periodic returns the periodic effect.
// (https://wiki.libsdl.org/SDL_HapticPeriodic)
func (he HapticEffect) Periodic() *HapticPeriodic {
	return (*HapticPeriodic)(unsafe.Pointer(&he[0]))
}

// Condition returns the condition effect.
// (https://wiki.libsdl.org/SDL_HapticCondition)
func (he HapticEffect) Condition() *HapticCondition {
	return (*HapticCondition)(unsafe.Pointer(&he[0]))
}

// Ramp returns the ramp effect.
// (https://wiki.libsdl.org/SDL_HapticRamp)
func (he HapticEffect) Ramp() *HapticRamp {
	return (*HapticRamp)(unsafe.Pointer(&he[0]))
}

// LeftRight returns the left/right effect.
// (https://wiki.libsdl.org/SDL_HapticLeftRight)
func (he HapticEffect) LeftRight() *HapticLeftRight {
	return (*HapticLeftRight)(unsafe.Pointer(&he[0]))
}

// Custom returns the custom effect.
// (https://wiki.libsdl.org/SDL_HapticCustom)
func (he HapticEffect) Custom() *HapticCustom {
	return (*HapticCustom)(unsafe.Pointer(&he[0]))
}

// SetType sets the happtic effect type.
// (https://wiki.libsdl.org/SDL_HapticEffect)
func (he HapticEffect) SetType(typ uint16) {
	*((*uint16)(unsafe.Pointer(&he[0]))) = typ
}

func (h *Haptic) cptr() *C.SDL_Haptic {
	return (*C.SDL_Haptic)(unsafe.Pointer(h))
}

// NumHaptics returns the number of haptic devices attached to the system.
// (https://wiki.libsdl.org/SDL_NumHaptics)
func NumHaptics() int {
	return int(C.SDL_NumHaptics())
}

// HapticName returns the implementation dependent name of a haptic device.
// (https://wiki.libsdl.org/SDL_HapticName)
func HapticName(index int) string {
	return (C.GoString)(C.SDL_HapticName(C.int(index)))
}

// HapticOpen opens a haptic device for use.
// (https://wiki.libsdl.org/SDL_HapticOpen)
func HapticOpen(index int) *Haptic {
	return (*Haptic)(unsafe.Pointer(C.SDL_HapticOpen(C.int(index))))
}

// HapticOpened reports whether the haptic device at the designated index has been opened.
// (https://wiki.libsdl.org/SDL_HapticOpened)
func HapticOpened(index int) int {
	return int(C.SDL_HapticOpened(C.int(index)))
}

// HapticIndex returns the index of a haptic device.
// (https://wiki.libsdl.org/SDL_HapticIndex)
func HapticIndex(h *Haptic) int {
	return int(C.SDL_HapticIndex(h.cptr()))
}

// MouseIsHaptic reports whether or not the current mouse has haptic capabilities.
// (https://wiki.libsdl.org/SDL_MouseIsHaptic)
func MouseIsHaptic() int {
	return int(C.SDL_MouseIsHaptic())
}

// HapticOpenFromMouse open a haptic device from the current mouse.
// (https://wiki.libsdl.org/SDL_HapticOpenFromMouse)
func HapticOpenFromMouse() *Haptic {
	return (*Haptic)(unsafe.Pointer(C.SDL_HapticOpenFromMouse()))
}

// JoystickIsHaptic reports whether a joystick has haptic features.
// (https://wiki.libsdl.org/SDL_JoystickIsHaptic)
func JoystickIsHaptic(joy *Joystick) int {
	return int(C.SDL_JoystickIsHaptic(joy.cptr()))
}

// HapticOpenFromJoystick opens a haptic device for use from a joystick device.
// (https://wiki.libsdl.org/SDL_HapticOpenFromJoystick)
func HapticOpenFromJoystick(joy *Joystick) *Haptic {
	return (*Haptic)(unsafe.Pointer(C.SDL_HapticOpenFromJoystick(joy.cptr())))
}

// Close closes a haptic device previously opened with HapticOpen().
// (https://wiki.libsdl.org/SDL_HapticClose)
func (h *Haptic) Close() {
	C.SDL_HapticClose(h.cptr())
}

// NumAxes returns the number of haptic axes the device has.
// (https://wiki.libsdl.org/SDL_HapticNumAxes)
func (h *Haptic) NumAxes() int {
	return int(C.SDL_HapticNumAxes(h.cptr()))
}

// NumEffects returns the number of effects a haptic device can store.
// (https://wiki.libsdl.org/SDL_HapticNumEffects)
func (h *Haptic) NumEffects() int {
	return int(C.SDL_HapticNumEffects(h.cptr()))
}

// NumEffectsPlaying reutrns the number of effects a haptic device can play at the same time.
// (https://wiki.libsdl.org/SDL_HapticNumEffectsPlaying)
func (h *Haptic) NumEffectsPlaying() int {
	return int(C.SDL_HapticNumEffectsPlaying(h.cptr()))
}

// Query returns haptic device's supported features in bitwise manner.
// (https://wiki.libsdl.org/SDL_HapticQuery)
func (h *Haptic) Query() uint {
	return uint(C.SDL_HapticQuery(h.cptr()))
}

// EffectSupported reports whether an effect is supported by a haptic device.
// (https://wiki.libsdl.org/SDL_HapticEffectSupported)
func (h *Haptic) EffectSupported(he *HapticEffect) int {
	_he := (*C.SDL_HapticEffect)(unsafe.Pointer(he))
	return int(C.SDL_HapticEffectSupported(h.cptr(), _he))
}

// NewEffect creates a new haptic effect on a specified device.
// (https://wiki.libsdl.org/SDL_HapticNewEffect)
func (h *Haptic) NewEffect(he *HapticEffect) int {
	_he := (*C.SDL_HapticEffect)(unsafe.Pointer(he))
	return int(C.SDL_HapticNewEffect(h.cptr(), _he))
}

// UpdateEffect updates the properties of an effect.
// (https://wiki.libsdl.org/SDL_HapticUpdateEffect)
func (h *Haptic) UpdateEffect(effect int, data *HapticEffect) int {
	_data := (*C.SDL_HapticEffect)(unsafe.Pointer(data))
	return int(C.SDL_HapticUpdateEffect(h.cptr(), C.int(effect), _data))
}

// RunEffect runs the haptic effect on its associated haptic device.
// (https://wiki.libsdl.org/SDL_HapticRunEffect)
func (h *Haptic) RunEffect(effect int, iterations uint32) int {
	return int(C.SDL_HapticRunEffect(h.cptr(), C.int(effect), C.Uint32(iterations)))
}

// StopEffect stops the haptic effect on its associated haptic device.
// (https://wiki.libsdl.org/SDL_HapticStopEffect)
func (h *Haptic) StopEffect(effect int) int {
	return int(C.SDL_HapticStopEffect(h.cptr(), C.int(effect)))
}

// DestroyEffect destroys a haptic effect on the device.
// (https://wiki.libsdl.org/SDL_HapticDestroyEffect)
func (h *Haptic) DestroyEffect(effect int) {
	C.SDL_HapticDestroyEffect(h.cptr(), C.int(effect))
}

// GetEffectStatus returns the status of the current effect on the specified haptic device.
// (https://wiki.libsdl.org/SDL_HapticGetEffectStatus)
func (h *Haptic) GetEffectStatus(effect int) int {
	return int(C.SDL_HapticGetEffectStatus(h.cptr(), C.int(effect)))
}

// SetGain sets the global gain of the specified haptic device.
// (https://wiki.libsdl.org/SDL_HapticSetGain)
func (h *Haptic) SetGain(gain int) int {
	return int(C.SDL_HapticSetGain(h.cptr(), C.int(gain)))
}

// SetAutocenter sets the global autocenter of the device.
// (https://wiki.libsdl.org/SDL_HapticSetAutocenter)
func (h *Haptic) SetAutocenter(autocenter int) int {
	return int(C.SDL_HapticSetAutocenter(h.cptr(), C.int(autocenter)))
}

// Pause pauses a haptic device.
// (https://wiki.libsdl.org/SDL_HapticPause)
func (h *Haptic) Pause() int {
	return int(C.SDL_HapticPause(h.cptr()))
}

// Unpause unpauses a haptic device.
// (https://wiki.libsdl.org/SDL_HapticUnpause)
func (h *Haptic) Unpause() int {
	return int(C.SDL_HapticUnpause(h.cptr()))
}

// StopAll stops all the currently playing effects on a haptic device.
// (https://wiki.libsdl.org/SDL_HapticStopAll)
func (h *Haptic) StopAll() int {
	return int(C.SDL_HapticStopAll(h.cptr()))
}

// RumbleSupported reports whether rumble is supported on a haptic device.
// (https://wiki.libsdl.org/SDL_HapticRumbleSupported)
func (h *Haptic) RumbleSupported() int {
	return int(C.SDL_HapticRumbleSupported(h.cptr()))
}

// RumbleInit initializes the haptic device for simple rumble playback.
// (https://wiki.libsdl.org/SDL_HapticRumbleInit)
func (h *Haptic) RumbleInit() int {
	return int(C.SDL_HapticRumbleInit(h.cptr()))
}

// RumblePlay runs a simple rumble effect on a haptic device.
// (https://wiki.libsdl.org/SDL_HapticRumblePlay)
func (h *Haptic) RumblePlay(strength float32, length uint32) int {
	return int(C.SDL_HapticRumblePlay(h.cptr(), C.float(strength), C.Uint32(length)))
}

// RumbleStop stops the simple rumble on a haptic device.
// (https://wiki.libsdl.org/SDL_HapticRumbleStop)
func (h *Haptic) RumbleStop() int {
	return int(C.SDL_HapticRumbleStop(h.cptr()))
}
