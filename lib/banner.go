package lib

import "fmt"

func Banner() {
	banner := "" +
		"\t\b\b\b\b\b██████╗ ██╗   ██╗███████╗██╗  ██╗███████╗██████╗ \n" +
		"\t\b\b\b\b\b██╔══██╗██║   ██║██╔════╝██║  ██║██╔════╝██╔══██╗\n" +
		"\t\b\b\b\b\b██████╔╝██║   ██║███████╗███████║█████╗  ██████╔╝\n" +
		"\t\b\b\b\b\b██╔═══╝ ██║   ██║╚════██║██╔══██║██╔══╝  ██╔══██╗\n" +
		"\t\b\b\b\b\b██║     ╚██████╔╝███████║██║  ██║███████╗██║  ██║\n" +
		"\t\b\b\b\b\b╚═╝      ╚═════╝ ╚══════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝"
	sub := "" +
		"Used for various security intelligence and article push\n" +
		"\tAuthor: Bingan\tLicense: Multiple"
	fmt.Println(banner)
	fmt.Println(sub)
}
