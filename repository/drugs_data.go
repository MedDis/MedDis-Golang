package repository

import (
	"context"
	"fmt"
	"gsc_rest/model"
	"strconv"
)

func GetAllDrugsData() []model.DrugsComposition {
	// fmt.Println("Get all drugs composition")
	ctx := context.Background()

	var result []model.DrugsComposition

	// var code string = strconv.Itoa(rand.Intn(999999-100000) + 100000)

	rows, err := DB.QueryContext(ctx, "select * from drugs_tb")

	if err != nil {
		// fmt.Println(err.Error())
		// return 0, "", err
	}

	var id, containAlcohol int
	var name, img, desc string

	for rows.Next() {

		if err := rows.Scan(&id, &name, &img, &containAlcohol, &desc); err != nil {
			// fmt.Println(err.Error())
		}

		// fmt.Println(strconv.Itoa(id) + ", " + strconv.Itoa(containAlcohol) + ", " + name + ", " + img + ", " + desc)

		ctx := context.Background()
		rows2, err := DB.QueryContext(ctx, "select country, synonim from drugs_synonim where drugs_id ="+strconv.Itoa(id))

		if err != nil {
			fmt.Println(err.Error())
			// return 0, "", err
		}

		var country, synonim string
		var synonimList []model.Synonim

		for rows2.Next() {
			if err := rows2.Scan(&country, &synonim); err != nil {
				// fmt.Println(err.Error())
			}

			// fmt.Println(country + ", " + synonim)
			synonimList = append(synonimList, model.Synonim{
				Country: country,
				Name:    synonim,
			})
		}

		ctx2 := context.Background()
		rows3, err := DB.QueryContext(ctx2, "select effects from drugs_effect where drugs_id ="+strconv.Itoa(id))

		if err != nil {
			fmt.Println(err.Error())
			// return 0, "", err
		}

		var effects string
		var effectList []string

		for rows3.Next() {
			if err := rows3.Scan(&effects); err != nil {
				// fmt.Println(err.Error())
			}

			// fmt.Println(effects)
			effectList = append(effectList, effects)
		}

		ctx3 := context.Background()
		rows4, err := DB.QueryContext(ctx3, "select `status`, min_age, max_age, `weight_condition`, `max_consumption`, cpd from drugs_dosage where drugs_id ="+strconv.Itoa(id))

		if err != nil {
			fmt.Println(err.Error())
			// return 0, "", err
		}

		var min_age, max_age, weight_condition, cpd int
		var status, max_consumption string
		var dosageList []model.Dosage

		for rows4.Next() {
			if err := rows4.Scan(&status, &min_age, &max_age, &weight_condition, &max_consumption, &cpd); err != nil {
				// fmt.Println(err.Error())
			}

			// fmt.Printf("stat : %v, minage : %v, maxage : %v, weight : %v, maxcons : %v, cpd : %v\n", status, min_age, max_age, weight_condition, max_consumption, cpd)

			dosageList = append(dosageList, model.Dosage{
				Type: status,
				Age: model.Ages{
					MinAge: min_age,
					MaxAge: max_age,
				},
				WeightCondition: weight_condition,
				MaxConsumption:  max_consumption,
				Cpd:             cpd,
			})
		}
		// return code, time, nil
		result = append(result, model.DrugsComposition{
			Name:    name,
			Synonim: synonimList,
			Img:     img,
			Dosages: dosageList,
			Effect:  effectList,
			Alcohol: containAlcohol == 1,
		})
	}

	return result
}

func GetDrugsByID(idDrugs int) (model.DrugsComposition, bool) {
	ctx := context.Background()

	var result model.DrugsComposition
	var notEmpty bool = false

	// var code string = strconv.Itoa(rand.Intn(999999-100000) + 100000)

	rows, err := DB.QueryContext(ctx, "select * from drugs_tb where id ="+strconv.Itoa(idDrugs))

	if err != nil {
		// fmt.Println(err.Error())
		// return 0, "", err
	}

	var id, containAlcohol int
	var name, img, desc string

	for rows.Next() {
		if err := rows.Scan(&id, &name, &img, &containAlcohol, &desc); err != nil {
			// fmt.Println(err.Error())
		}

		// fmt.Println(strconv.Itoa(id) + ", " + strconv.Itoa(containAlcohol) + ", " + name + ", " + img + ", " + desc)

		ctx := context.Background()
		rows2, err := DB.QueryContext(ctx, "select country, synonim from drugs_synonim where drugs_id ="+strconv.Itoa(id))

		if err != nil {
			fmt.Println(err.Error())
			// return 0, "", err
		}

		var country, synonim string
		var synonimList []model.Synonim

		for rows2.Next() {
			if err := rows2.Scan(&country, &synonim); err != nil {
				// fmt.Println(err.Error())
			}

			// fmt.Println(country + ", " + synonim)
			synonimList = append(synonimList, model.Synonim{
				Country: country,
				Name:    synonim,
			})
		}

		ctx2 := context.Background()
		rows3, err := DB.QueryContext(ctx2, "select effects from drugs_effect where drugs_id ="+strconv.Itoa(id))

		if err != nil {
			fmt.Println(err.Error())
			// return 0, "", err
		}

		var effects string
		var effectList []string

		for rows3.Next() {
			if err := rows3.Scan(&effects); err != nil {
				// fmt.Println(err.Error())
			}

			// fmt.Println(effects)
			effectList = append(effectList, effects)
		}

		ctx3 := context.Background()
		rows4, err := DB.QueryContext(ctx3, "select `status`, min_age, max_age, `weight_condition`, `max_consumption`, cpd from drugs_dosage where drugs_id ="+strconv.Itoa(id))

		if err != nil {
			fmt.Println(err.Error())
			// return 0, "", err
		}

		var min_age, max_age, weight_condition, cpd int
		var status, max_consumption string
		var dosageList []model.Dosage

		for rows4.Next() {
			if err := rows4.Scan(&status, &min_age, &max_age, &weight_condition, &max_consumption, &cpd); err != nil {
				// fmt.Println(err.Error())
			}

			// fmt.Printf("stat : %v, minage : %v, maxage : %v, weight : %v, maxcons : %v, cpd : %v\n", status, min_age, max_age, weight_condition, max_consumption, cpd)

			dosageList = append(dosageList, model.Dosage{
				Type: status,
				Age: model.Ages{
					MinAge: min_age,
					MaxAge: max_age,
				},
				WeightCondition: weight_condition,
				MaxConsumption:  max_consumption,
				Cpd:             cpd,
			})
		}
		// return code, time, nil

		result = model.DrugsComposition{
			Name:    name,
			Synonim: synonimList,
			Img:     img,
			Dosages: dosageList,
			Effect:  effectList,
			Alcohol: containAlcohol == 1,
		}
		notEmpty = true
	}

	return result, notEmpty
}
