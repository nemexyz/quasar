package controller

import (
	"io"
	"net/http"
	"quasar/api/models"
	"quasar/domain"
	"quasar/repository"

	"github.com/gin-gonic/gin"
)

var repo = repository.NewSatelliteRepository()

func getEnemies(satellites []domain.Satellite) (domain.AllianceResponse, error) {
	alliance := domain.NewAlliance(satellites)
	location, err := alliance.FindEnemyLocation()
	if err != nil {
		return domain.AllianceResponse{}, err
	}
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

func PostTopSecretSplit(c *gin.Context) {
	var body models.TSSRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repo.SaveMessage(c.Param("satellite_name"), body.Distance, body.Message); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Satellite not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Message saved"})
	}
}

func GetTopSecretSplit(c *gin.Context) {
	satellites, err := repo.GetAllSatellitesWithLastMessages()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	dereferencedSatellites := make([]domain.Satellite, len(satellites))
	for i, satellite := range satellites {
		dereferencedSatellites[i] = *satellite
	}
	result, err := getEnemies(dereferencedSatellites)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
