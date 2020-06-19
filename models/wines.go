package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

// Marshal Wine model to ORM
type Wine struct {
	Id          int       `orm:"auto"`
	Product     string    `orm:"size(64)"`
	Description string    `orm:"size(128)"`
	Price       float32   `orm:"null"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime);null"`
}

func GetAllWines() ([]*Wine, error) {

	o := orm.NewOrm()

	var wines []*Wine
	qs := o.QueryTable("wine")
	_, err := qs.OrderBy("Id").All(&wines)

	if err != nil {
		return nil, err
	}

	return wines, nil
}

func GetWine(id int) (*Wine, error) {

	o := orm.NewOrm()

	// Fetch wine by Id
	wine := Wine{Id: id}
	err := o.Read(&wine)
	if err != nil {
		return nil, err
	}

	return &wine, nil
}

func AddWine(wine Wine) (*Wine, error) {

	o := orm.NewOrm()

	id, err := o.Insert(&wine)
	if err != nil {
		return nil, err
	}

	return GetWine(int(id))
}

func UpdateWine(uid string, wine Wine) (*Wine, error) {

	o := orm.NewOrm()

	// Assign Id to update
	id, err := strconv.Atoi(uid)
	if err != nil {
		return nil, err
	}
	wine.Id = id

	var fields []string

	// Update changed fields
	if wine.Description != "" {
		fields = append(fields, "Description")
	}

	if wine.Product != "" {
		fields = append(fields, "Product")
	}

	if wine.Price != 0.0 {
		fields = append(fields, "Price")
	}

	_, err = o.Update(&wine, fields...)
	if err != nil {
		return nil, err
	}

	// Return JSON for update
	return GetWine(id)
}

func DeleteWine(id int) (*Wine, error) {

	// wine := new(Wine)
	o := orm.NewOrm()

	// Select the object to delete
	wine := Wine{Id: id}
	err := o.Read(&wine)
	if err != nil {
		return nil, err
	}

	// delete
	_, err = o.Delete(&wine)

	return &wine, nil
}
