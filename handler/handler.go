package handler

import (
	"strconv"

	"github.com/labstack/echo"
)

// struct
type kelas struct {
	Nama string `json:"nama"` // awalin dengan huruf kapital
	Umur int    `json:"umur"`
}

// Bar - Ini pasngannya Foo
func Bar(c echo.Context) error {
	param := c.Param("param")

	input := c.QueryParam("input")

	toNumber, err := strconv.Atoi(input)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, echo.Map{
		"input":  input,
		"number": toNumber,
		"param":  param,
	})
}

// PostSomething - Ini untuk posting something
func PostSomething(c echo.Context) error {
	k := new(kelas)

	if err := c.Bind(k); err != nil { // ngeBind itu ngerubah aga bisa dibaca , ngebind itu fitur dr echo
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, k)

}

// CreateData - Ini untuk masukin data ke table_fahmis
func CreateData(c echo.Context) error {
	inputData := new(TableFahmi)

	if err := c.Bind(inputData); err != nil { //
		return echo.NewHTTPError(500, err.Error())
	}

	db, err := getConnection()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	defer db.Close()

	err = inputData.InsertData(db)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(201, inputData)
}

// GetData - INi untuk get data by its name
func GetData(c echo.Context) error {
	// Ini untuk Get input user dari query param, dengan key 'name'
	name := c.QueryParam("nama")

	// Ini initiate data dengan type TableFahmi
	data := TableFahmi{
		Nama: name,
	}

	// Get koneksi ke database untuk di parsing saat ingin query
	db, err := getConnection()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	// Saat fungsi sudah beres, close connection
	defer db.Close()

	// Query data ke database menggunakan method dari si "TableFahmi"
	response, err := data.GetDataByName(db)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, response)
}

func GetAllData(c echo.Context) error {
	var d TableFahmi

	db, err := getConnection()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	defer db.Close()

	res, err := d.GetAllData(db)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, res)
}
