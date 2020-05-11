package app

import (
	"golang-microservices/src/api/controllers/polo"
	"golang-microservices/src/api/controllers/repositories"
)

func mapUrl() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
