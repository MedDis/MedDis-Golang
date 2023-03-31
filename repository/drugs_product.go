package repository

import (
	"context"
	"fmt"
	"gsc_rest/model"
	"strconv"
)

func GetAllDrugProductsData(pagination int) []model.DrugsProductMini {
	ctx := context.Background()

	var result []model.DrugsProductMini

	// var code string = strconv.Itoa(rand.Intn(999999-100000) + 100000)
	rows, err := DB.QueryContext(ctx, "SELECT `id`, `name` FROM drugs_product LIMIT "+strconv.Itoa((pagination*10)-10)+", "+strconv.Itoa(pagination*10))

	if err != nil {
		fmt.Println(err.Error())
		// return 0, "", err
	}

	var name string
	var id int

	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			// fmt.Println(err.Error())
		}

		result = append(result, model.DrugsProductMini{
			Id:   id,
			Name: name,
		})
	}
	// print(result)
	return result
}

func GetDrugProductsDataById(id int) (model.DrugsProduct, bool) {
	ctx := context.Background()

	var result model.DrugsProduct
	var isEmpty bool = false

	// var code string = strconv.Itoa(rand.Intn(999999-100000) + 100000)
	rows, err := DB.QueryContext(ctx, "SELECT `kat`, `name`, `deskripsi`, `indikasi`, `komposisi`, `dosis`, `aturan_pakai`, `no_registrasi`, `efek`, `manufaktur` FROM drugs_product where id = "+strconv.Itoa(id))

	if err != nil {
		fmt.Println(err.Error())
		// return 0, "", err
	}

	var kat, name, desc, indication, composition, dosage, rule, regist, effect, manufacture string

	for rows.Next() {
		if err := rows.Scan(&kat, &name, &desc, &indication, &composition, &dosage, &rule, &regist, &effect, &manufacture); err != nil {
			// fmt.Println(err.Error())
		}

		result = model.DrugsProduct{
			Categories:  kat,
			Name:        name,
			Indication:  indication,
			Composition: composition,
			Dosage:      dosage,
			UseRule:     rule,
			Desc:        desc,
			NoRegist:    regist,
			Effect:      effect,
			Manufacture: manufacture,
		}
		isEmpty = true
	}
	// print(result)
	return result, isEmpty
}
