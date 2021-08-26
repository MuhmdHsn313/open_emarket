package settings

import "gorm.io/gorm"

const SECRET_KEY = "@kWYGZ@CAxS\\b(&wx.&xtx@gr}2c\\aD-"

const CONFIG_VALUE_VALIDATOR = `([A-Z]+)\s?=\s?([\w\d.\\\/ ]+)`

const CONFIG_FILE_PATH = "settings.config"

var DB *gorm.DB
