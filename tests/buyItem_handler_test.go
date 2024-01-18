package tests

import (
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/freshusername/to-buy-api/api"
	"github.com/freshusername/to-buy-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type BuyItemsTestSuite struct {
	suite.Suite
	App *fiber.App
}

func TestBuyItemsTestSuite(t *testing.T) {
	//t.Setenv("DB_USER", "#dbuser#")
	err := godotenv.Load("../.env.local")
	if err != nil {
		log.Fatal("Problem loading .env.local file")
		os.Exit(-1)
	}
	suite.Run(t, new(BuyItemsTestSuite))
}

func (suite *BuyItemsTestSuite) SetupTest() {
	app := fiber.New()
	app.Get("/buyItems", api.HandleGetBuyItems)
	go func() {
		app.Listen(":3000")
	}()
	suite.App = app
}

func (suite *BuyItemsTestSuite) TestHandleGetBuyItemss() {
	database.ConnectDb()

	req := httptest.NewRequest("GET", "/buyItems", nil)
	res, err := suite.App.Test(req)
	suite.Nil(err)
	suite.Equal(200, res.StatusCode)
}
