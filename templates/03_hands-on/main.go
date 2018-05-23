package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name    string
	Address string
}

type zipcode struct {
	Code   uint16
	Hotels []hotel
}

type city struct {
	Name     string
	ZipCodes []zipcode
}

type region struct {
	Name   string
	Cities []city
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	regions := []region{
		region{
			Name: "Northern",
			Cities: []city{
				city{
					Name: "Test City",
					ZipCodes: []zipcode{
						zipcode{
							Code: 11234,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel",
									Address: "Somewhere",
								},
								hotel{
									Name:    "Test Hotel 2",
									Address: "Here!",
								},
							},
						},
						zipcode{
							Code: 11111,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel 3",
									Address: "Somewhere Else",
								},
								hotel{
									Name:    "Test Hotel 4",
									Address: "There!",
								},
							},
						},
					},
				},
				city{
					Name: "Test City 2",
					ZipCodes: []zipcode{
						zipcode{
							Code: 22222,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel",
									Address: "Somewhere",
								},
								hotel{
									Name:    "Test Hotel 2",
									Address: "Here!",
								},
							},
						},
						zipcode{
							Code: 33333,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel 3",
									Address: "Somewhere Else",
								},
								hotel{
									Name:    "Test Hotel 4",
									Address: "There!",
								},
							},
						},
					},
				},
			},
		},
		region{
			Name: "Central",
			Cities: []city{
				city{
					Name: "Test City",
					ZipCodes: []zipcode{
						zipcode{
							Code: 11234,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel",
									Address: "Somewhere",
								},
								hotel{
									Name:    "Test Hotel 2",
									Address: "Here!",
								},
							},
						},
						zipcode{
							Code: 11111,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel 3",
									Address: "Somewhere Else",
								},
								hotel{
									Name:    "Test Hotel 4",
									Address: "There!",
								},
							},
						},
					},
				},
				city{
					Name: "Test City 2",
					ZipCodes: []zipcode{
						zipcode{
							Code: 22222,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel",
									Address: "Somewhere",
								},
								hotel{
									Name:    "Test Hotel 2",
									Address: "Here!",
								},
							},
						},
						zipcode{
							Code: 33333,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel 3",
									Address: "Somewhere Else",
								},
								hotel{
									Name:    "Test Hotel 4",
									Address: "There!",
								},
							},
						},
					},
				},
			},
		},
		region{
			Name: "Southern",
			Cities: []city{
				city{
					Name: "Test City",
					ZipCodes: []zipcode{
						zipcode{
							Code: 11234,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel",
									Address: "Somewhere",
								},
								hotel{
									Name:    "Test Hotel 2",
									Address: "Here!",
								},
							},
						},
						zipcode{
							Code: 11111,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel 3",
									Address: "Somewhere Else",
								},
								hotel{
									Name:    "Test Hotel 4",
									Address: "There!",
								},
							},
						},
					},
				},
				city{
					Name: "Test City 2",
					ZipCodes: []zipcode{
						zipcode{
							Code: 22222,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel",
									Address: "Somewhere",
								},
								hotel{
									Name:    "Test Hotel 2",
									Address: "Here!",
								},
							},
						},
						zipcode{
							Code: 33333,
							Hotels: []hotel{
								hotel{
									Name:    "Test Hotel 3",
									Address: "Somewhere Else",
								},
								hotel{
									Name:    "Test Hotel 4",
									Address: "There!",
								},
							},
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, regions)
	if err != nil {
		log.Fatalln(err)
	}
}
