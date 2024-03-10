package domain

import (
	"errors"
	"math"
	"strings"
)

type Alliance struct {
	satellites []Satellite
}

func NewAlliance(satellites []Satellite) *Alliance {
	return &Alliance{
		satellites: satellites,
	}
}

func (af *Alliance) Decode() string {
	decoded := []string{}
	originaLength := math.MaxInt64
	for _, s := range af.satellites {
		length := len(s.message)
		if length < originaLength {
			originaLength = length
		}
	}
	for _, s := range af.satellites {
		s.FixMsgDelay(originaLength)
	}
	if len(af.satellites) > 0 {
		f := af.satellites[0]
		r := af.satellites[1:]
		for i := 0; i < originaLength; i++ {
			word := f.message[i]
			if word != "" {
				decoded = append(decoded, word)
			} else {
				valid := ""
				for _, s := range r {
					word2 := s.message[i]
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
	if len(af.satellites) == 1 {
		return [2]float64{}, errors.New("Not enough active satellites to find enemy location")
	}
	if len(af.satellites) == 2 {
		fstSat := af.satellites[0]
		sndSat := af.satellites[1]
		int := fstSat.IntersectionWith(&sndSat)
		if af.AreSame(int[0], int[1]) {
			return int[0], nil
		}
		return [2]float64{}, errors.New("Not enough active satellites to find enemy location")
	}
	fstThreeSats := af.satellites[:3]
	kenobi := fstThreeSats[0]
	skywalker := fstThreeSats[1]
	sato := fstThreeSats[2]
	intKenSky := kenobi.IntersectionWith(&skywalker)
	intKenSato := kenobi.IntersectionWith(&sato)
	intSkySato := skywalker.IntersectionWith(&sato)
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
		return [2]float64{}, errors.New("Not enough active satellites to find enemy location")
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
