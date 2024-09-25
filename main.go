package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)
type Company struct{
	Name string 
	Address string
	LogoLocation string
}
type Ticket struct{
	ID int
	ShowName string
	ShowTime string
	Language           string
    ShowVenue          string
    SeatNumber         string
    Cost               float64
    Screen             string
    TicketCount        int
    ShowPosterLocation string
}
func getPageFooter() core.Row{
	 return row.New(2).Add(
		col.New(12).Add(
			text.New("Powered by Blaise Golang ticketing System",props.Text{
				Style: fontstyle.Italic,
				Size: 8,
				Align: align.Center,
				Color: &props.Color{Red: 255, Green: 120, Blue: 218},

			}),
		),
	 )
}
func pageHeader(c Company)  core.Row{
	return row.New(16).Add(
		image.NewFromFileCol(4,c.LogoLocation,props.Rect{
			Center: false,
			Percent: 100,
		}),
		col.New(2),
		col.New(6).Add(
			text.New(c.Name,props.Text{
				Style:fontstyle.Bold ,
				Size: 10,
			}),
			text.New(c.Address,props.Text{
				Top: 6,
				Size: 10,
			}),
		),

	
	)
}
func getMaroto(c Company, t Ticket) core.Maroto{
	cfg:=config.NewBuilder().WithDimensions(120,200).Build()
	mrt:=maroto.New(cfg)
	err:=mrt.RegisterHeader(pageHeader(c))
	if err!=nil{
		log.Println("Error registering")
	}
	mrt.AddRow(6)
	mrt.AddRow(4,line.NewCol(12,props.Line{
		Thickness: 0.2,
		Color: &props.Color{Red:200,Green: 200,Blue: 200},
	}))
	mrt.AddRow(6)
	mrt.AddRows(getShowDetails(t)...)
	mrt.AddRow(8)
	err=mrt.RegisterFooter(getPageFooter())
	if err!=nil{
		log.Println("errro registering the footer")
	}
	
	return mrt
}
func getShowDetails(t Ticket) []core.Row{
	rows:=[]core.Row{
		row.New(30).Add(
			image.NewFromFileCol(5,t.ShowPosterLocation,props.Rect{
				Center: true,
				Percent: 80,
				Left: 1,
			}),
		 col.New(7).Add(
			text.New(t.Language,props.Text{
			Style: fontstyle.Bold,
			Size: 10,
		}),
		text.New(t.Language,props.Text{
			Top:6,
			Size: 8,
			Style: fontstyle.Normal,
			Color: &props.Color{Red: 95, Green: 95, Blue: 95},
		}),
		text.New(t.ShowTime,props.Text{
			Top: 12,
			Style: fontstyle.Bold,
			Size:  10,
		}),
		text.New(t.ShowVenue, props.Text{
			Top:   18,
			Style: fontstyle.Normal,
			Size:  8,
			Color: &props.Color{Red: 95, Green: 95, Blue: 95},
		}),
)	,
		),
		row.New(6),
		row.New(1).Add(
			line.NewCol(12,props.Line{
				Thickness: 0.2,
				Color: &props.Color{Red: 200, Green: 200, Blue: 200},
				SizePercent: 100,
				Style: linestyle.Dashed,
			}),
		),
		row.New(3),
		row.New(16).Add(
			col.New(2).Add(
				text.New(strconv.Itoa(t.TicketCount),props.Text{
					Style: fontstyle.Bold,
					Size: 24,
					Align: align.Center,
				}),
				text.New("Tickets",props.Text{
					
					Top:   12,
                    Style: fontstyle.Normal,
                    Size:  8,
                    Color: &props.Color{Red: 95, Green: 95, Blue: 95},
                    Align: align.Center,
				})),
				col.New(2),
				col.New(8).Add(
					text.New(t.Screen,props.Text{
						Size: 8,
						Color: &props.Color{Red: 95, Green: 95, Blue: 95},
					}),
					text.New(t.SeatNumber, props.Text{
						Top:   6,
						Style: fontstyle.Bold,
						Size:  14,
					}),

				),
			),
		row.New(3),
		row.New(1).Add(
            line.NewCol(12, props.Line{
                Thickness:   0.2,
                Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
                SizePercent: 100,
                Style:       linestyle.Dashed,
            }),
        ),
		row.New(6),
		row.New(20).Add(
			code.NewQrCol(12,
				fmt.Sprintf("%v\n %v \n %v \n %v",t.ID, t.ShowName, t.ShowTime, t.ShowVenue),props.Rect{
					Center: true,
					Percent: 100,
				},
			),
			
		),
		row.New(10).Add(
            col.New(12).Add(text.New(fmt.Sprintf("Booking ID: %v", t.ID), props.Text{
                Style: fontstyle.Normal,
                Size:  8,
                Align: align.Center,
                Top:   2,
            })),
        ),
		row.New(1).Add(
            line.NewCol(12, props.Line{
                Thickness:   0.2,
                Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
                SizePercent: 100,
                Style:       linestyle.Solid,
            }),
        ),
		row.New(3),
		row.New(10).Add(
			code.NewBarCol(12,strconv.Itoa(t.ID),
		props.Barcode{
			Center: true,
			Percent: 100,
		}),
		),
	
	}


	return rows

}
func main(){
	c:=Company{
		Name: "ShowBees Ticketing",
		Address: "1234 Main St,City,State 123",
		LogoLocation:"./logo.jpeg",
	}
	t:=Ticket{
		ID:                 1,
        ShowName:           "Planet of the Gophers: The War Begins",
        ShowTime:           "Sat 01/01/2022 7:00 PM",
        Language:           "English",
        ShowVenue:          "Gophedorium",
        SeatNumber:         "Platinum - A1, A2",
        Cost:               620.00,
        Screen:             "Screen 1",
        TicketCount:        2,
        ShowPosterLocation: "./event_location.png",
	}
	m:=getMaroto(c,t)
	document,err:=m.Generate()
	// 
	filename:=fmt.Sprintf("ticket-%d.pdf",t.ID)
	if err!=nil{
		log.Println("Error generating PDF")

	}
	// file error
 if _,err:=os.Stat("./images");os.IsNotExist(err){
err=os.Mkdir("images",0755)
if err!=nil{
	log.Println("Error creating directory")
} 

 }
 err=document.Save("./images/"+filename)
if err!=nil{
	log.Println("unable to save the file",err)
}


}