package repositories

import (
	"database/sql"
	"fmt"
	"github.com/sanzharanarbay/repository-service-pattern/application/models"
	"log"
	"time"
)

type ItemRepository struct {
	dbClient *sql.DB
}

func NewItemRepository(dbClient *sql.DB) *ItemRepository {
	return &ItemRepository{
		dbClient: dbClient,
	}
}

type ItemRepositoryInterface interface {
	GetItemById(ID int) (*models.Item, error)
	GetAllItems() ([]*models.Item, error)
	SaveItem(*models.Item) (bool, error)
	DeleteItem(ID int) (bool, error)
	UpdateItem(*models.Item) (bool, error)
}

func (i *ItemRepository) GetItemById(ID int) (*models.Item, error) {
	var item models.Item
	err := i.dbClient.QueryRow(`SELECT * FROM items WHERE id=$1`, ID).Scan(&item.ID, &item.Name, &item.Cost, &item.Description,
		&item.CreatedAt, &item.UpdatedAt)
	switch err {
	case sql.ErrNoRows:
		log.Printf("No rows were returned!")
		return nil, nil
	case nil:
		return &item, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return &item, nil
}

func (i *ItemRepository) GetAllItems() (*[]models.Item, error) {
	rows, err := i.dbClient.Query("SELECT * FROM items")
	if err != nil {
		fmt.Printf("ERROR SELECT QUERY - %s", err)
		return nil, err
	}
	var itemList []models.Item
	for rows.Next() {
		var item models.Item
		err = rows.Scan(&item.ID, &item.Name, &item.Cost, &item.Description,
			&item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			fmt.Printf("ERROR QUERY SCAN - %s", err)
			return nil, err
		}
		itemList = append(itemList, item)
	}
	return &itemList, nil
}

func (i *ItemRepository) SaveItem(item *models.Item) (bool, error) {
	sqlStatement := `INSERT into items (name,cost, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := i.dbClient.Exec(sqlStatement, item.Name, item.Cost, item.Description, time.Now().Local(), time.Now().Local())
	if err != nil {
		log.Printf("ERROR EXEC INSERT QUERY - %s", err)
		return false, err
	}
	return true, nil
}

func (i *ItemRepository) DeleteItem(ID int) (bool, error) {
	_, err := i.dbClient.Exec(`DELETE FROM items WHERE id=$1`, ID)
	if err != nil {
		log.Printf("ERROR EXEC DELETE QUERY - %s", err)
		return false, err
	}
	return true, nil
}

func (i *ItemRepository) UpdateItem(item *models.Item, ItemID int) (bool, error) {
	sqlStatement := `UPDATE items SET name=$1, cost=$2, description=$3, updated_at=$4 WHERE id=$5`
	_, err := i.dbClient.Exec(sqlStatement, item.Name, item.Cost, item.Description, time.Now().Local(), ItemID)
	if err != nil {
		fmt.Printf("ERROR EXEC UPDATE QUERY - %s", err)
	}
	return true, nil
}
