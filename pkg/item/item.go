package item

import (
	"errors"
	"fmt"
	"net/http"
	"netlog/model"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func SaveItems(c *gin.Context) {
	var params struct {
		Url string `json:"url"`
	}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if ok := isExist(params.Url); ok {
		c.JSON(200, gin.H{"message": "Article already parsed"})
		return
	}

	if err := parseItems(params.Url); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Success"})
}

func parseItems(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("status code error: " + res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}
	doc.Find("article p a").Each(func(i int, s *goquery.Selection) {

		item := model.Item{Title: doc.Find("title").Text(), OriginUrl: url, Name: s.Text(), Href: s.AttrOr("href", "")}
		model.DB.Create(&item)
	})
	return nil
}

func isExist(url string) bool {
	var item model.Item
	result := model.DB.Where("origin_url = ?", url).First(&item)
	return result.RowsAffected == 1
}

func GetItems(c *gin.Context) {
	var page string = c.Query("page")
	var pageSize string = c.Query("pageSize")
	var keyword string = c.Query("keyword")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		pageNum = 1
	}
	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeNum = 10
	}
	fmt.Print(page, pageSize, keyword)

	var items []model.Item
	model.DB.Where("name LIKE ?", "%"+keyword+"%").
		Offset((pageNum - 1) * pageSizeNum).
		Limit(pageSizeNum).
		Find(&items)
	c.JSON(200, items)

}
