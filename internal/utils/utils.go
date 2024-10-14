package utils

import "homework/web/internal/models"

func NewItems() []models.Item {
	return make([]models.Item, 0)
}

func AddItem(items *[]models.Item, i *models.Item) {
	*items = append(*items, *i)
}

func SearchByCaption(items []models.Item, caption string) *[]models.Item {
	matchItem := NewItems()
	for i := range items {
		if items[i].Caption == caption {
			AddItem(&matchItem, &items[i])
		}
	}
	return &matchItem
}
