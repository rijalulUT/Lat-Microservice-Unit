package unit

import (
	"math/rand"
	"time"
	"unitmicroservice/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.opencensus.io/trace"
)

type Units struct {
	DB *gorm.DB
}

// type TblUnit struct {
// 	Id       int    `json:"id"`
// 	NamaUnit string `json:"nama_unit"`
// 	KodeUnit string `json:"kode_unit"`
// }

func (p *Units) GetUnit(c *gin.Context) {
	db := p.DB
	var unit []model.TblUnit

	db.Find(&unit)

	c.JSON(200, gin.H{
		"data": unit,
	})

	serviceb(c)
}

func (p *Units) GetUnitById(c *gin.Context) {
	var unit model.TblUnit
	db := p.DB
	id := c.Param("id")

	db.Where("id = ?", id).Find(&unit)

	c.JSON(200, gin.H{
		"nama_unit": unit.NamaUnit,
		"kode_unit": unit.KodeUnit,
	})
}

func serviceb(c *gin.Context) {
	_, span := trace.StartSpan(c, "/products")
	defer span.End()
	time.Sleep(time.Duration(rand.Intn(800)+200) * time.Millisecond)
}
