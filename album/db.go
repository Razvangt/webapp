package album

import (
	"webapp/common"
)

func all() []Album {
	db := common.GetDB()
	var albums []Album
	db.Find(&albums)
	return albums
}

// Add the new album to the slice
func add(newAlbum *Album) {
	db := common.GetDB()
	db.Create(&newAlbum)
}

func find(id *string) (Album, error) {
	db := common.GetDB()
	var a Album
	result := db.Find(&a, id)

	if result.Error != nil {
		return Album{}, result.Error
	}
	return a, nil
}

func update(updated *Album) error {
	db := common.GetDB()
	result := db.Save(&updated)
	return result.Error
}

func delete(id *string) error {
	db := common.GetDB()
	result := db.Delete(&Album{}, id)
	return result.Error
}
