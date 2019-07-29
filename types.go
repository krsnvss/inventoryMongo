package main

import (
	"gopkg.in/mgo.v2/bson"
)

type baseBoard struct {
	Manufacturer string `json:"Manufacturer,omitempty"`
	Product      string `json:"Product,omitempty"`
	SerialNumber string `json:"SerialNumber,omitempty"`
	Version      string `json:"Version,omitempty"`
}
type bios struct {
	Manufacturer string `json:"Manufacturer,omitempty"`
	Name         string `json:"Name,omitempty"`
	ReleaseDate  string `json:"ReleaseDate,omitempty"`
	SerialNumber string `json:"SerialNumber,omitempty"`
	Version      string `json:"Version,omitempty"`
}
type computerSystem struct {
	CurrentTimeZone     string `json:"CurrentTimeZone,omitempty"`
	Domain              string `json:"Domain,omitempty"`
	Name                string `json:"Name,omitempty"`
	SystemFamily        string `json:"SystemFamily,omitempty"`
	SystemSKUNumber     string `json:"SystemSKUNumber,omitempty"`
	TotalPhysicalMemory string `json:"TotalPhysicalMemory,omitempty"`
	UserName            string `json:"UserName,omitempty"`
}
type cpu struct {
	Manufacturer  string `json:"Manufacturer,omitempty"`
	MaxClockSpeed string `json:"MaxClockSpeed,omitempty"`
	Name          string `json:"Name,omitempty"`
	NumberOfCores string `json:"NumberOfCores,omitempty"`
}
type diskDrive struct {
	MediaType string `json:"MediaType,omitempty"`
	Model     string `json:"Model,omitempty"`
	Size      string `json:"Size,omitempty"`
}
type display struct {
	Name                string `json:"Name,omitempty"`
	MonitorManufacturer string `json:"MonitorManufacturer,omitempty"`
	MonitorType         string `json:"MonitorType,omitempty"`
}
type electronics struct {
	Cartridge    string        `json:"Cartridge,omitempty"`
	ID           bson.ObjectId `bson:"_id,omitempty"`
	IP           string        `json:"IP,omitempty"`
	Dt           string        `json:"dt"`
	Manufacturer string        `json:"Manufacturer,omitempty"`
	Model        string        `json:"Model,omitempty"`
	Name         string        `json:"Name,omitempty"`
	Type         string        `json:"Type,omitempty"`
	SerialNumber string        `json:"SerialNumber,omitempty"`
	Printouts    string        `json:"Printouts,omitempty"`
	Purchase     string        `json:"Purchase,omitempty"`
	Repairs      []repair      `json:"Repairs,omitempty"`
}
type keyboard struct {
	Description string `json:"Description,omitempty"`
	DeviceID    string `json:"DeviceID,omitempty"`
	Name        string `json:"Name,omitempty"`
}
type mouse struct {
	Description  string `json:"Description,omitempty"`
	DeviceID     string `json:"DeviceID,omitempty"`
	Manufacturer string `json:"Manufacturer,omitempty"`
	Name         string `json:"Name,omitempty"`
}
type networkDevice struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	IP           string        `json:"IP,omitempty"`
	Dt           string        `json:"dt"`
	Firmware     string        `json:"Firmware,omitempty"`
	MAC          string        `json:"Mac,omitempty"`
	Manufacturer string        `json:"Manufacturer,omitempty"`
	Model        string        `json:"Model,omitempty"`
	Name         string        `json:"Name,omitempty"`
	Type         string        `json:"Type,omitempty"`
	SerialNumber string        `json:"SerialNumber,omitempty"`
	Purchase     string        `json:"Purchase,omitempty"`
	Uptime       string        `json:"Uptime,omitempty"`
}
type nic struct {
	AdapterType  string `json:"AdapterType,omitempty"`
	MACAddress   string `json:"MACAddress,omitempty"`
	Manufacturer string `json:"Manufacturer,omitempty"`
	Name         string `json:"Name,omitempty"`
	ProductName  string `json:"ProductName,omitempty"`
	Speed        string `json:"Speed,omitempty"`
}
type os struct {
	CountryCode  string `json:"CountryCode,omitempty"`
	Manufacturer string `json:"Manufacturer,omitempty"`
	Caption      string `json:"Caption,omitempty"`
	InstallDate  string `json:"InstallDate,omitempty"`
	OSLanguage   string `json:"OSLanguage,omitempty"`
	Version      string `json:"Version,omitempty"`
}
type pclist []sysinfo
type printer struct {
	Caption    string `json:"Caption,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	PortName   string `json:"PortName,omitempty"`
}
type repair struct {
	Dt      string `json:"dt"`
	Price   string `json:"price,omitempty"`
	Comment string `json:"comment,omitempty"`
}
type sysinfo struct {
	ID             bson.ObjectId     `bson:"_id,omitempty"`
	Bios           bios              `json:"bios"`
	BaseBoard      baseBoard         `json:"baseboard"`
	CPU            []cpu             `json:"cpu"`
	ComputerSystem computerSystem    `json:"computersystem"`
	DiskDrive      []diskDrive       `json:"diskdrive"`
	Display        []display         `json:"display"`
	Dt             string            `json:"dt"`
	Keyboard       []keyboard        `json:"keyboard"`
	Mouse          []mouse           `json:"mouse"`
	NIC            []nic             `json:"nic"`
	OS             os                `json:"os"`
	Printer        []printer         `json:"printer"`
	Tz             tz                `json:"tz"`
	User           string            `json:"user"`
	Video          []videoController `json:"video"`
}
type tz struct {
	Caption string `json:"Caption,omitempty"`
}
type videoController struct {
	Caption                     string `json:"Caption,omitempty"`
	CurrentHorizontalResolution string `json:"CurrentHorizontalResolution,omitempty"`
	CurrentVerticalResolution   string `json:"CurrentVerticalResolution,omitempty"`
	CurrentBitsPerPixel         string `json:"CurrentBitsPerPixel,omitempty"`
	VideoProcessor              string `json:"VideoProcessor,omitempty"`
}
