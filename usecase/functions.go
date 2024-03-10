package usecase

import (
	"fmt"
	"quasar/domain"
)

func GetLocation(distKen, distSky, distSato float64) [2]float64 {
	kenobiCoord := [2]float64{-500, -200}
	skywalkerCoord := [2]float64{100, -100}
	satoCoord := [2]float64{500, 100}

	kenobi := domain.NewSatellite("kenobi", kenobiCoord)
	kenobi.ReceiveMessage(distKen, nil)

	skywalker := domain.NewSatellite("skywalker", skywalkerCoord)
	skywalker.ReceiveMessage(distSky, nil)

	sato := domain.NewSatellite("sato", satoCoord)
	sato.ReceiveMessage(distSato, nil)

	result := domain.NewAlliance([]domain.Satellite{*kenobi, *skywalker, *sato})

	location, err := result.FindEnemyLocation()
	if err != nil {
		fmt.Println("Error finding enemy location:", err)
		return [2]float64{}
	}
	return location
}

func GetMessage(messages [][]string) string {
	kenobi := domain.NewSatellite("Kenobi", [2]float64{0, 0})
	kenobi.ReceiveMessage(0, messages[0])

	skywalker := domain.NewSatellite("Skywalker", [2]float64{0, 0})
	skywalker.ReceiveMessage(0, messages[1])

	sato := domain.NewSatellite("Sato", [2]float64{0, 0})
	sato.ReceiveMessage(0, messages[2])

	result := domain.NewAlliance([]domain.Satellite{*kenobi, *skywalker, *sato})

	return result.Decode()
}
