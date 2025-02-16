package main

import (
	"encoding/json"
	"fmt"

	arrowhead "github.com/johankristianss/arrowhead/pkg/arrowhead"
	"github.com/johankristianss/arrowhead/pkg/rpc"
)

type Car struct {
	Brand string `json:"brand"`
	Color string `json:"color"`
}

type InMemoryCarRepository struct {
	cars []Car
}

type CreateCarService struct {
	inMemoryCarRepository *InMemoryCarRepository
}

func (s *CreateCarService) HandleRequest(params *arrowhead.Params) ([]byte, error) {
	fmt.Println("CreateCarService called, creating car")
	car := Car{}
	err := json.Unmarshal(params.Payload, &car)
	if err != nil {
		return nil, err
	}
	fmt.Println("Car: ", car)
	s.inMemoryCarRepository.cars = append(s.inMemoryCarRepository.cars, car)
	return nil, nil
}

type GetCarService struct {
	inMemoryCarRepository *InMemoryCarRepository
}

func (s *GetCarService) HandleRequest(params *arrowhead.Params) ([]byte, error) {
	carsJSON, err := json.Marshal(s.inMemoryCarRepository.cars)
	if err != nil {
		return nil, err
	}
	return carsJSON, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	framework, err := arrowhead.CreateFramework()
	checkError(err)

	inMemoryCarRepository := &InMemoryCarRepository{}
	createCarService := &CreateCarService{inMemoryCarRepository: inMemoryCarRepository}
	getCarService := &GetCarService{inMemoryCarRepository: inMemoryCarRepository}

	framework.HandleService(createCarService, rpc.POST, "create-car", "/carfactory")
	framework.HandleService(getCarService, rpc.GET, "get-car", "/carfactory")

	err = framework.ServeForever()
	checkError(err)
}
