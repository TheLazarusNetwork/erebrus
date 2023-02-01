package dbinit

import (
	"log"

	"github.com/TheLazarusNetwork/erebrus/api/v1/flowid"
	"github.com/TheLazarusNetwork/erebrus/dbconfig"
)

func Init() error {
	db := dbconfig.GetDb()
	err := db.AutoMigrate(&flowid.FlowId{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
