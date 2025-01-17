package nasType

// DNN 9.11.2.1A
// DNN Row, sBit, len = [0, 0], 8 , INF
type DNN struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewDNN(iei uint8) (dNN *DNN) {
	dNN = &DNN{}
	dNN.SetIei(iei)
	return dNN
}

// DNN 9.11.2.1A
// Iei Row, sBit, len = [], 8, 8
func (a *DNN) GetIei() (iei uint8) {
	return a.Iei
}

// DNN 9.11.2.1A
// Iei Row, sBit, len = [], 8, 8
func (a *DNN) SetIei(iei uint8) {
	a.Iei = iei
}

// DNN 9.11.2.1A
// Len Row, sBit, len = [], 8, 8
func (a *DNN) GetLen() (len uint8) {
	return a.Len
}

// DNN 9.11.2.1A
// Len Row, sBit, len = [], 8, 8
func (a *DNN) SetLen(len uint8) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

// DNN 9.11.2.1A
// DNN Row, sBit, len = [0, 0], 8 , INF
func (a *DNN) GetDNN() (dNN []uint8) {
	dNN = make([]uint8, len(a.Buffer))
	copy(dNN, a.Buffer)
	return dNN
}

// DNN 9.11.2.1A
// DNN Row, sBit, len = [0, 0], 8 , INF
func (a *DNN) SetDNN(dNN []uint8) {
	copy(a.Buffer, dNN)
}
