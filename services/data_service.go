package dataService

import (
	"luck-go/daos/data_dao"
	"luck-go/dtos/data"
)

func FetchData() ([]data.DataDTO, error) {
	dataList, err := data_dao.GetData()
	if err != nil {
		return nil, err
	}

	var dataDTOs []data.DataDTO
	for _, d := range dataList {
		dataDTOs = append(dataDTOs, data.DataDTO{
			ID:    d.ID,
			Name:  d.Name,
			Value: d.Value,
		})
	}

	return dataDTOs, nil
}
