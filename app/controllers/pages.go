package controllers

import (
	//"fmt"
	//"os"
	//"bufio"
	//"//io/ioutil"
	//"html/template"
	"path/filepath"
	//"strings"

	"github.com/revel/revel"
	//"github.com/russross/blackfriday"
	//"gopkg.in/yaml.v2"
	//"github.com/pksunkara/pygments"
)


type CurrPage struct {
	//Title string
	Version string
	SectionUrl string
	SectionTitle string
	PageUrl string
	PageTitle string
	//Version string
	Lang string
}

//var Site *SiteStruct

func GetCurrPage(section, section_title, version, lang, page string) CurrPage {

	s := CurrPage{SectionUrl: section, SectionTitle: section_title, PageUrl: page, Version: version, Lang: lang}
	return s
}



type Pages struct {
	*revel.Controller
}

// /robots.txt - Only allow spiders on prod site
func (c Pages) RobotsTxt() revel.Result {

	txt := "User-agent: *\n"
	if revel.Config.BoolDefault("site.live", false)  == false {
		txt += "Disallow: /\n"
	}
	txt += "\n"

	return c.RenderText(txt)
}

// main home page
func (c Pages) Index() revel.Result {
	return c.Render()
}



// render an expected markdown file
func (c Pages) Markdown(site_section, ver, lang, page string) revel.Result {



	cPage := GetCurrPage(site_section, "Manual", ver, lang, page)

	nav := GetNav(site_section)
	c.RenderArgs["nav"] = nav


	page_no_ext := page
	if filepath.Ext(page) == ".html" { // wtf includes the .
		page_no_ext = page[0: len(page) - 5]
	}

	// use template.HTML to "unescape" encoding.. ie proper html not &lt;escaped
	pdata := ReadMarkdownPage(site_section, page_no_ext)
	c.RenderArgs["page_content"] = pdata.HTML
	cPage.PageTitle = pdata.Title


	c.RenderArgs["cPage"] = cPage

	return c.Render()
}

func (c Pages) Godoc(go_file string) revel.Result {



	cPage := GetCurrPage("docs/godoc", "Go Docs", "0.16", "en", go_file)

	cPage.PageTitle = go_file


	c.RenderArgs["cPage"] = cPage

	return c.Render()
}



func (c Pages) Github() revel.Result {
	return c.Render()
}

