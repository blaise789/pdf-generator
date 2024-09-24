package main

import (
	"fmt"
	"log"
	"os"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
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
func pageHeader( c Company)core.Row{
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
	return mrt
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
        ShowPosterLocation: "./poster.png",
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
err=document.Save("./images/"+filename)
if err!=nil{
	log.Println("unable to save the file",err)
}
 }


}