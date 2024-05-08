package infra

import (
	local "credits_manager/error"
	"encoding/csv"
	"os"
)

func SaveCSV(filepath string, data [][]string) error {
	file, err := os.Create(filepath)
	if err != nil {
		local.HandleErrorMessage("Cant create csv file", err)
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			local.HandleErrorMessage("Cant close csv file", err)
		}
	}()

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(data); err != nil {
		local.HandleErrorMessage("Cant write data to csv file", err)
		return err
	}
	writer.Flush()
	if writer.Error() != nil {
		local.HandleErrorMessage("Cant flush csv file", writer.Error())
		return writer.Error()
	}

	return nil
}
