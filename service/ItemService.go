package service

import (
	"myshop/model"
	"myshop/repo"
)

// AddItem AddItem
func AddItem(item *model.Item) error {
	return repo.AddItem(item)
}
