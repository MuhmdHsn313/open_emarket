package settings

import "gorm.io/gorm"

const SECRET_KEY = "@kWYGZ@CAxS\\b(&wx.&xtx@gr}2c\\aD-"

const CONFIG_VALUE_VALIDATOR = `([A-Z]+)\s?=\s?([\w\d.\\\/ ]+)`

const CONFIG_FILE_PATH = "settings.config"

var DB *gorm.DB

const HTML_GLOB_PATH = "template/*"

const STATIC_PATH = "static/"

const MEDIA_PATH = "media/"

const MaxMultipartMemory = 8 << 20

const JWT_EXPIRES_HOURS = 72

const JWT_LOGIN_SUBJECT = "LOGIN_EVENT"
