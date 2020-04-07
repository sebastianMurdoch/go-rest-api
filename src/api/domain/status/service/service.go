package service

import "github.com/sebastianMurdoch/go-rest-api/src/api/domain/status"

type StatusServiceImpl struct {

}

func NewItemsServiceImpl() *StatusServiceImpl {
	return &StatusServiceImpl{}
}

func (ss *StatusServiceImpl) Get() (*status.Status, error){
	return &status.Status{Value:"Ok"}, nil
}