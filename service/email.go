package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Medical Discovery Support<meddiscov@gmail.com>"
const CONFIG_AUTH_EMAIL = "<yourmail>"
const CONFIG_AUTH_PASSWORD = "<yourCode>"

func SendCodeVerification(clientMail string, code int) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", clientMail)
	mailer.SetHeader("Subject", "Kode Verifikasi MedDis")
	mailer.SetBody("text/html", htmlBodyHandler(clientMail, code))

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	// log.Println("Mail sent!")
}

func htmlBodyHandler(clientEmail string, code int) string {
	var data string = `
	<html>
	<head>
		<style>
			body {
				font-family: arial;
				font-size: 10pt;
			}
	
			img {
				display: block;
				width: 200px;
				margin-right: auto;
				margin-left: auto;
				margin-bottom: 5%;
			}

			h1 {
				color: #008080;
				text-align: center;
			}

			.div-box {
				background-color: #ffffff;
			}

			h2 {
				font-size: 32px;
			}
	
			.parent {
				padding: 15px 15px;
				margin: auto;
				max-width: 500px;
			}
	
			.container {
				padding: 20px;
			}

			.gap {
				height: 20px;
			}
	
			.center {
				margin: auto;
				width: 30%;
				padding: 10px;
			}
	
			@media screen and (max-width: 400px) {
				.parent-container {
					width: 400px;
					margin: 20px auto;
				}
	
				p .box {
					padding-top: 2%;
				}
			}
		</style>
	</head>
	
	<body>
		<div style="background-color: #F1F1F1;">
			<div class="parent">
				<div class="div-box container">
					<h1>MedDis</h1>
					<div class="gap"></div>
					<img src="https://bit.ly/3lw8JO8" width height="50%" alt="Logo">
					<div class="gap"></div>
					<h1>Hi, ` + strings.Split(clientEmail, "@")[0] + `</h1>
					<div> 
						<p align="center">Thanks for register and welcome to MedDis!</p>
						<p align="center">We want to make sure it's really you. Please enter the following verification code</p>
						<div class="center" align="center">
							<h2>` + strconv.Itoa(code) + `</h2>
						</div>
						<div class="gap"></div>
						<p class="end"><i>This is an automated email. Do not reply this email.</i></p>
						<p>© Code in 2023</p>
					</div>
				</div>
			</div>
		</div>
	</body>
	</html>`
	fmt.Println(strconv.Itoa(code))
	return data
}
