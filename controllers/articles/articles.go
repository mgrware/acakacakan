package articles


import (
  "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
  "net/http"
  "github.com/gin-gonic/gin"
  "mongo/models"

)

func Index(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	articles := []models.Article{}
	err := db.C(models.CollectionArticle).Find(nil).Sort("-updated_on").All(&articles)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{"articles": articles})
}

func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	article := models.Article{}
	err := c.Bind(&article)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionArticle).Insert(article)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{"articles": article})
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	article := models.Article{}
	err := c.Bind(&article)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err = db.C(models.CollectionArticle).Update(query, article)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{"articles": article})
}

func Edit(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	article := models.Article{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionArticle).FindId(oID).One(&article)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{"articles": article})

}
