package repository

import (
	"context"
	"fmt"
	"gsc_rest/service"
	"math/rand"
	"strconv"
	"strings"
)

func AddNewCodeVerification(email string) (string, string) {
	ctx := context.Background()

	var code string = strconv.Itoa(rand.Intn(999999-100000) + 100000)

	_, err := DB.ExecContext(ctx, "insert into code_verif values (null, \""+email+"\", "+code+", now())")

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			rows, err := DB.QueryContext(ctx, "select code, time FROM code_verif where email=\""+email+"\"")

			if err != nil {
				fmt.Println(err.Error())
			}

			var code, time string

			for rows.Next() {

				if err := rows.Scan(&code, &time); err != nil {
					fmt.Println(err.Error())
				}
				// fmt.Println(code, time)
				return code, time
			}

			rows.Close()
		}
	}
	return code, service.CurrentUTime()
}

func UpdateCodeVerification(email string) (int, bool) {

	ctx := context.Background()

	var code int = rand.Intn(999999-100000) + 100000

	if _, err := DB.ExecContext(ctx, "UPDATE code_verif set code = "+strconv.Itoa(code)+", time = now() WHERE email=\""+email+"\""); err != nil {
		fmt.Println(err.Error())
		return code, true
	}
	return code, false
}

func SelectCodeVerification(email string) (int, string, error) {

	ctx := context.Background()

	rows, err := DB.QueryContext(ctx, "SELECT code, time FROM code_verif WHERE email=\""+email+"\"")

	if err != nil {
		// fmt.Println(err.Error())
		return 0, "", err
	}

	var code int
	var time string

	for rows.Next() {

		if err := rows.Scan(&code, &time); err != nil {
			fmt.Println(err.Error())
		}
		// fmt.Println(code, time)
		return code, time, nil
	}

	return code, time, nil
}

func RemoveCodeVerification(email string) {

	ctx := context.Background()

	if _, err := DB.ExecContext(ctx, "delete from code_verif WHERE email=\""+email+"\""); err != nil {
		fmt.Println(err.Error())
	}
}
