package quests

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Difficulty [actsCount]Act

//NewDifficulty returns Difficulty from packed bytes
func NewDifficulty(r io.Reader) (Difficulty, error) {
	d := Difficulty{}
	var err error
	for i := range d {
		d[i], err = NewAct(r, ActId(i))
		if err != nil {
			return d, err
		}
	}
	return d, nil
}

//GetPacked returns packed Difficulty into []byte
func (d *Difficulty) GetPacked() ([]byte, error) {
	var buf bytes.Buffer

	for _, act := range d {
		if err := binary.Write(&buf, binaryEndian, act.GetPacked()); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}