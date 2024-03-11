package controller

import (
	"fmt"
	"io"
	"net/http"
	"quasar/domain"
	"quasar/repository"

	"github.com/gin-gonic/gin"
)

var repo = repository.NewSatelliteRepository()

func getEnemies(satellites []domain.Satellite) (domain.AllianceResponse, error) {
	alliance := domain.NewAlliance(satellites)
	fmt.Println("Alliance:", alliance)
	location, err := alliance.FindEnemyLocation()
	if err != nil {
		return domain.AllianceResponse{}, err
	}
	fmt.Println("Location:", location)
	message := alliance.Decode()
	return domain.AllianceResponse{
		Position: domain.PositionResponse{
			X: int64(location[0]),
			Y: int64(location[1]),
		},
		Message: message,
	}, nil
}

func PostTopSecret(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input, err := domain.UnmarshalAlliance(bodyBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list := []domain.Satellite{}
	for _, s := range input.Satellites {
		sc, err := repo.GetSatelliteByName(s.Name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		sc.ReceiveMessage(s.Distance, s.Message)
		list = append(list, *sc)
	}

	result, err := getEnemies(list)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
