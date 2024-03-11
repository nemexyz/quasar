package domain

import (
	"encoding/json"
	"errors"
	"math"
	"strings"
)

type Alliance struct {
	Satellites []Satellite `json:"satellites"`
}

func UnmarshalAlliance(data []byte) (Alliance, error) {
	var r Alliance
	err := json.Unmarshal(data, &r)
	return r, err
}

func NewAlliance(satellites []Satellite) *Alliance {
	return &Alliance{
		Satellites: satellites,
	}
}

func (af *Alliance) Decode() string {
	decoded := []string{}
	originaLength := math.MaxInt64
	for _, s := range af.Satellites {
		length := len(s.Message)
		if length < originaLength {
			originaLength = length
		}
	}
	for _, s := range af.Satellites {
		s.FixMsgDelay(originaLength)
	}
	if len(af.Satellites) > 0 {
		f := af.Satellites[0]
		r := af.Satellites[1:]
		for i := 0; i < originaLength; i++ {
			word := f.Message[i]
			if word != "" {
				decoded = append(decoded, word)
			} else {
				valid := ""
				for _, s := range r {
					word2 := s.Message[i]
					if word2 != "" {
						valid = word2
					}
				}
				decoded = append(decoded, valid)
			}
		}
	}
	return strings.Join(decoded, ", ")
}

func (af *Alliance) FindEnemyLocation() ([2]float64, error) {
	if len(af.Satellites) == 1 {
		return [2]float64{}, errors.New("not enough active satellites to find enemy location")
	}
	if len(af.Satellites) == 2 {
		fstSat := af.Satellites[0]
		sndSat := af.Satellites[1]
		int, err := fstSat.IntersectionWith(&sndSat)
		if err != nil {
			return [2]float64{}, err
		}
		if af.AreSame(int[0], int[1]) {
			return int[0], nil
		}
		return [2]float64{}, errors.New("not enough active satellites to find enemy location")
	}
	fstThreeSats := af.Satellites[:3]
	kenobi := fstThreeSats[0]
	skywalker := fstThreeSats[1]
	sato := fstThreeSats[2]
	intKenSky, err := kenobi.IntersectionWith(&skywalker)
	intKenSato, err2 := kenobi.IntersectionWith(&sato)
	intSkySato, err3 := skywalker.IntersectionWith(&sato)
	if err != nil || err2 != nil || err3 != nil {
		return [2]float64{}, errors.New("not enough active satellites to find enemy location")
	}
	enemyCoord := [2]float64{}
	for _, int := range intKenSky {
		isInKenSato := false
		isInSkySato := false
		for _, i := range intKenSato {
			if af.AreEqual(i, int) {
				isInKenSato = true
				break
			}
		}
		for _, i := range intSkySato {
			if af.AreEqual(i, int) {
				isInSkySato = true
				break
			}
		}
		if isInKenSato && isInSkySato {
			enemyCoord = int
			break
		}
	}
	if enemyCoord == [2]float64{} {
		return [2]float64{}, errors.New("not enough active satellites to find enemy location")
	}
	return enemyCoord, nil
}

func (af *Alliance) AreSame(pointA, pointB [2]float64) bool {
	return pointA[0] == pointB[0] && pointA[1] == pointB[1]
}

func (af *Alliance) AreEqual(pointA, pointB [2]float64) bool {
	return af.CompareWithError(pointA[0], pointB[0]) && af.CompareWithError(pointA[1], pointB[1])
}

func (af *Alliance) CompareWithError(numA, numB float64) bool {
	return math.Abs(math.Abs(numA)-math.Abs(numB)) <= 0.04
}
