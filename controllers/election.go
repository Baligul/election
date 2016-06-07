package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _ "github.com/lib/pq"
	models "github.com/Baligul/election/models"
)

func init() {
    orm.RegisterDriver("postgres", orm.DRPostgres)
    orm.RegisterDataBase("default", "postgres", "postgres://ggxssikrsehequ:sQElIpN-CHqcFFNAx7mJO31Y3v@ec2-54-225-93-34.compute-1.amazonaws.com:5432/da6obv8tnlvcev")
    orm.RegisterModel(new(models.Voter))
}

type ElectionController struct {
	beego.Controller
}

func (e *ElectionController) Home() {
	e.Layout = "basic-layout.tpl"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["Header"] = "header.tpl"
	e.LayoutSections["Footer"] = "footer.tpl"
	e.TplName = "election/home.tpl"
}

/*
func (e *ElectionController) Update() {
	o := orm.NewOrm()
	o.Using("default")
	flash := beego.NewFlash()

	// convert the string value to an int
	if articleId, err := strconv.Atoi(e.Ctx.Input.Param(":id")); err == nil {
		article := models.Article{Id: articleId}
		// attempt to load the record from the database
		if o.Read(&article) == nil {
			// set the Client and Url properties to arbitrary values
			article.Client = "Sitepoint"
			article.Url = "http://www.google.com"
			if num, err := o.Update(&article); err == nil {
				flash.Notice("Record Was Updated.")
				flash.Store(&e.Controller)
				beego.Info("Record Was Updated. ", num)
			}
		} else {
			flash.Notice("Record Was NOT Updated.")
			flash.Store(&e.Controller)
			beego.Error("Couldn't find article matching id: ", articleId)
		}
	} else {
		flash.Notice("Record Was NOT Updated.")
		flash.Store(&e.Controller)
		beego.Error("Couldn't convert id from a string to a number. ", err)
	}

	// redirect afterwards
	e.Redirect("/e/view", 302)
}*/

func (e *ElectionController) View() {
	e.Layout = "basic-layout.tpl"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["Header"] = "header.tpl"
	e.LayoutSections["Footer"] = "footer.tpl"
	e.TplName = "election/view.tpl"

	flash := beego.ReadFromRequest(&e.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		e.Data["errors"] = ok
	}

	if ok := flash.Data["notice"]; ok != "" {
		// Display error messages
		e.Data["notices"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")
    
    var voters []*models.Voter
	num, err := o.QueryTable("voter").Filter("Section_name_english__exact","GHER MOHABBAT KHA").All(&voters)

	if err != orm.ErrNoRows && num > 0 {
		e.Data["records"] = voters
        e.Data["counts"] = num
	} else {
        panic(err)
    }
    /*

    db, err := sql.Open("postgres", "user=member password= dbname=election sslmode=disable")
    if err != nil {
        panic(err)
    }
    
    // Query
    rows, err := db.Query("SELECT * FROM voter")
    if err != nil {
        panic(err)
    }
    
	var voters []*models.Voter
    
    for rows.Next() {
        var Name string
        var RelationName string
        var Gender string
        var District string
        var AC string
        var SectionName string
        err = rows.Scan(&NameEnglish, &RelationNameEnglish, &Gender, &DistrictNameEnglish, &ACNameEnglish, &SectionNameEnglish)
        if err != nil {
            panic(err)
        }
        
        
        
        voters.append(voters,)
    }*/
    
    //e.Data["records"] = voters
}