package main

import (
	server "github.com/uncleTen276/attendance.git/server/app"
)

//	func main() {
//		err := godotenv.Load()
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		client, err := ethclient.Dial(os.Getenv("INFURA_URL"))
//		if err != nil {
//			log.Fatal(err)
//		}
//		contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
//		contract, err := attendance_record.NewAttendanceRecord(contractAddress, client)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		chaiId, err := client.NetworkID(context.Background())
//		if err != nil {
//			log.Fatal("chainId", err.Error())
//		}
//		fmt.Println(chaiId)
//
//		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chaiId)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		app := fiber.New()
//		app.Post("/attendance_record", func(c *fiber.Ctx) error {
//			tx, err := contract.RecordAtttendance(auth, big.NewInt(1), "test", "test", "tower A")
//			if err != nil {
//				return c.SendString(err.Error())
//			}
//			return c.SendString(tx.Hash().Hex())
//		})
//
//		app.Get("/attendance", func(c *fiber.Ctx) error {
//			tx, err := contract.GetAttendanceByEmployee(nil, big.NewInt(1))
//			if err != nil {
//				return c.SendString(err.Error())
//			}
//			return c.JSON(tx)
//		})
//
//		app.Put("/attendance", func(c *fiber.Ctx) error {
//			tx, err := contract.UpdateAttendance(auth, big.NewInt(1), "2024", "2025", "nothing here", "tower A")
//			if err != nil {
//				return c.SendString(err.Error())
//			}
//			return c.JSON(tx)
//		})
//		app.Listen(":8080")
//	}
//
//

// @title Fiber Go API
// @version 1.0
// @description greendeco
// @contact.name Nguyen Tri
// @contact.email tringuyen2762001@gmail.com
// @termsOfService
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	server.Serve()
}
