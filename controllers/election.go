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
    //orm.RegisterDataBase("default", "postgres", "user=member dbname=election sslmode=disable")
    orm.RegisterModel(new(models.Voter))
}

type ElectionController struct {
	beego.Controller
}

func 
(e *ElectionController) Home() {
	e.Layout = "basic-layout.tpl"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["Header"] = "header.tpl"
	e.LayoutSections["Footer"] = "footer.tpl"
	e.TplName = "election/home.tpl"
    
    var (
        votersCount int64
        muslimVotersCount int64
        otherVotersCount int64
        maleVotersCount int64
        femaleVotersCount int64
        err interface{}
    )
    
    o := orm.NewOrm()
	o.Using("default")
    qs := o.QueryTable("voter")
    
    flash := beego.ReadFromRequest(&e.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		e.Data["errors"] = ok
	}

	if ok := flash.Data["notice"]; ok != "" {
		// Display error messages
		e.Data["notices"] = ok
	}
    
    votersCount, err = qs.Count()
    
    if err != orm.ErrNoRows && votersCount > 0 {
        e.Data["VotersCount"] = votersCount
	} else {
        e.Data["VotersCount"] = "No voters found"
    }
    
    muslimVotersCount, err = qs.Filter("Religion_english__exact", "Muslim").Count()
    
    if err != orm.ErrNoRows && muslimVotersCount > 0 {
        e.Data["MuslimVotersCount"] = muslimVotersCount
	} else {
        e.Data["MuslimVotersCount"] = "No muslim voters found"
    }
    
    otherVotersCount, err = qs.Filter("Religion_english__exact", "Other").Count()
    
    if err != orm.ErrNoRows && otherVotersCount > 0 {
        e.Data["OtherVotersCount"] = otherVotersCount
	} else {
        e.Data["OtherVotersCount"] = "No other voters found"
    }   
    
    maleVotersCount, err = qs.Filter("Gender__exact", "M").Count()
    
    if err != orm.ErrNoRows && maleVotersCount > 0 {
        e.Data["MaleVotersCount"] = maleVotersCount
	} else {
        e.Data["MaleVotersCount"] = "No male voters found"
    }   
    
    femaleVotersCount, err = qs.Filter("Gender__exact", "F").Count()
    
    if err != orm.ErrNoRows && femaleVotersCount > 0 {
        e.Data["FemaleVotersCount"] = femaleVotersCount
	} else {
        e.Data["FemaleVotersCount"] = "No female voters found"
    }      
    
    e.Data["MuslimVotersPercentage"] = (float64(muslimVotersCount) / float64(votersCount)) * 100
    e.Data["OtherVotersPercentage"] = (float64(otherVotersCount) / float64(votersCount)) * 100
    e.Data["MaleVotersPercentage"] = (float64(maleVotersCount) / float64(votersCount)) * 100
    e.Data["FemaleVotersPercentage"] = (float64(femaleVotersCount) / float64(votersCount)) * 100
}

func (e *ElectionController) View() {
	e.Layout = "basic-layout.tpl"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["Header"] = "header.tpl"
	e.LayoutSections["Footer"] = "footer.tpl"
	e.TplName = "election/view.tpl"
    
    var (
        num int64
        votersCount int64
        criteriaVotersCount int64
        voters []*models.Voter
        err interface{}
    )
    
    o := orm.NewOrm()
	o.Using("default")
    qs := o.QueryTable("voter")
    
	flash := beego.ReadFromRequest(&e.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		e.Data["errors"] = ok
	}

	if ok := flash.Data["notice"]; ok != "" {
		// Display error messages
		e.Data["notices"] = ok
	}
    
    gender := e.GetString("gender")
    
    if gender == "male" {
        qs = qs.Filter("Gender__exact", "M")
    }
    
    if gender == "female" {
        qs = qs.Filter("Gender__exact", "F")
    }
    
    acname := e.GetString("acname")
    if acname != "" {
        qs = qs.Filter("Ac_name_english__exact", acname)
    }
    
    district := e.GetString("district")
    if district != "" {
        qs = qs.Filter("District_name_english__exact", district)
    }
    
    sectionname := e.GetString("sectionname")
    if sectionname != "" {
        qs = qs.Filter("Section_name_english__exact", sectionname)
    }
    
    religion := e.GetString("religion")
    
    if religion == "Muslim" {
        qs = qs.Filter("Religion_english__exact", "Muslim")
    }
    
    if religion == "Other" {
        qs = qs.Filter("Religion_english__exact", "Other")
    }
    
    num, err = qs.Limit(1000).All(&voters)
    criteriaVotersCount, err = qs.Count()
    
    voterOrm := orm.NewOrm()
	voterOrm.Using("default")
    votersCount, err = voterOrm.QueryTable("voter").Count()
    
    if err != orm.ErrNoRows && num > 0 {
		e.Data["records"] = voters
        e.Data["counts"] = num
        votersCountPercentage := (float64(criteriaVotersCount) / float64(votersCount)) * 100
        e.Data["CriteriaVotersCount"] = criteriaVotersCount
        e.Data["CriteriaVotersPercentage"] = votersCountPercentage
	} else {
        e.Data["counts"] = "No voters found with this criteria. Please try again!"
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

func (e *ElectionController) Form() {
	e.Layout = "basic-layout.tpl"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["Header"] = "header.tpl"
	e.LayoutSections["Footer"] = "footer.tpl"
	e.TplName = "election/view.tpl"
}