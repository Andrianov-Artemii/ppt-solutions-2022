package main

import "fmt"

func checkIt(username string, link string) string {
	return fmt.Sprintf("%s переоткрыл заявку на проверку, надо поглазеть👀! \n Ссылка на PR: %s", username, link)
}

func banAlert(username string, link string) string {
	return fmt.Sprintf("💢%s запушил решение тогда, когда этого делать было нельзя.💢 \n Ссылка на PR: %s", username, link)
}
