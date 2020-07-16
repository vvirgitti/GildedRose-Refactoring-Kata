package main

type Item struct {
	name            string
	sellIn, quality int
}

const agedBrie = "Aged Brie"
const backstage = "Backstage passes to a TAFKAL80ETC concert"
const sulfuras = "Sulfuras, Hand of Ragnaros"

func isAgedBrie(itemName string) bool {
	return itemName == agedBrie
}

func isBackstage(itemName string) bool {
	return itemName == backstage
}

func isSulfuras(itemName string) bool {
	return itemName == sulfuras
}

func decrease(value int) int {
	return value - 1
}

func increase(value int) int {
	return value + 1
}

func hasPositiveQuality(value int) bool {
	return value > 0
}

func decreaseSellin(item *Item) {
	if !isSulfuras(item.name) {
		item.sellIn = decrease(item.sellIn)
	}
}

func UpdateQuality(items []*Item) {
	for i, item := range items {
		itemName := item.name

		if !isAgedBrie(itemName) && !isBackstage(itemName) {
			if hasPositiveQuality(items[i].quality) && !isSulfuras(itemName) {
				items[i].quality =  decrease(items[i].quality)
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality =  increase(items[i].quality)
				if isBackstage(itemName) && items[i].sellIn < 11 && items[i].quality < 50 {
					items[i].quality = increase(items[i].quality)
					if items[i].sellIn < 6 && items[i].quality < 50 {
						items[i].quality = increase(items[i].quality)
					}
				}
			}
		}

		decreaseSellin(items[i])



		if items[i].sellIn < 0 {
			if isAgedBrie(itemName) && items[i].quality < 50 {
				items[i].quality =  increase(items[i].quality)
			}

			if !isAgedBrie(itemName) {
				if !isBackstage(itemName) &&  hasPositiveQuality(items[i].quality) && !isSulfuras(itemName) {
					items[i].quality =  decrease(items[i].quality)
				} else {
					items[i].quality = items[i].quality - items[i].quality
				}
			}
		}
	}
}
