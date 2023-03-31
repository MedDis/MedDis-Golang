package main

import (
	"gsc_rest/repository"
)

func InitService() {
	repository.DB = repository.GetConnection()
}
