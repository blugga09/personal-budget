package config

var Companies map[string]string = loadFile("./config/company_category.json")
var Categories map[string]string = loadFile("./config/categories.json")
